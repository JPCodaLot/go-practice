export class BaseElement extends HTMLElement {
  static createAttrHandler(context) {
    return {
      get(target, prop, receiver) {
        return context.shadowRoot.host.getAttribute(prop);
      },
      set(target, prop, value, receiver) {
        context[prop] = value;
        if (value) {
          context.shadowRoot.host.setAttribute(prop, value);
        } else {
          context.shadowRoot.host.removeAttribute(prop);
        }
        return true
      },
    }
  }

  constructor() {
    super();
  }

  connectedCallback() {
		this._internals = this.attachInternals();
    const shadow = this.attachShadow({ mode: "open" });
    this.attrs = new Proxy(this.constructor.attrs, BaseElement.createAttrHandler(this));
    shadow.innerHTML = this.render();
    shadow.adoptedStyleSheets = [this.constructor.styles];
  }
}

export function css(strings, ...values) {
  const sheet = new CSSStyleSheet();
  sheet.replaceSync(String.raw({ raw: strings }, ...values));
  return sheet;
}

export function html(strings, ...values) {
	return String.raw({ raw: strings }, ...values);
}

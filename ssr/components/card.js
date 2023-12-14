import {css, html, BaseElement} from './base.js'

// Create a class for the element
export class CardElement extends BaseElement {
  static observedAttributes = ["selected"];

  static attrs = {
    selected: "true",
  };

  static styles = css`
    :host {
      position: relative;
      display: block;
      border-radius: 4px;
      margin-bottom: 16px;
      max-width: 600px;
      box-shadow: 0 4px 8px 0 rgba(0,0,0,0.1);
      user-select: none;
      cursor: pointer;
      border: 1px solid #bbb;
      transition: box-shadow 200ms linear;
    }
    :host(:last-child) {
      margin-bottom: 0px;
    }
    :host([selected=true]) {
      outline: 2px solid #b9d0f5;
      border-color: #81a1d5;
    }

    :host > header {
      background-color: #eee;
      border-bottom: 1px solid #bbb;
      padding: 8px;
      border-radius: 4px 4px 0px 0px;
    }
    :host > div {
      padding: 16px 8px;
    }
    :host > footer {
      background-color: #eee;
      border-top: 1px solid #bbb;
      padding: 8px;
      border-radius: 0px 0px 4px 4px;
    }

    :host #status {
      position: absolute;
      top: 0;
      right: 0;
      margin: 8px;
      font-family: monospace;
    }
  `;

  constructor() {
    super();
		this.addEventListener("click", this.onClick);
  }

	onClick() {
    this.attrs.selected = !this.attrs.selected;
	}

  render() {
    return html`
      <header><slot name="header"></slot><span id="status"></span></header>
      <div><slot></slot></div>
      <footer><slot name="footer"></slot></footer>
    `;
  }

  attributeChangedCallback(name, _, value) {
    if (name === "selected") {
      const tag = this.shadowRoot.querySelector("#status");
      if (value) {
        tag.textContent = "Selected";
      } else {
        tag.textContent = "";
      }
    }
  }
}

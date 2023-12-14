export class SelectElement extends HTMLElement {
  static observedAttributes = ["multiselect"];
  constructor() {
    super();
	}
	connectedCallback() {
		this.items = this.querySelector('#items');
		this.template = this.querySelector('template');
	}
	get selected() {
	  return this.items.querySelectorAll('my-card[selected=true]');
	}
	add() {
    const item = this.template.content.cloneNode(true)
	  this.items.appendChild(item);
	}
	removeSelected() {
	  this.selected.forEach((item) => item.remove());
	}
  selectAll() {
	  const cards = this.items.querySelectorAll('my-card');
	  cards.forEach((item) => item.setAttribute("selected", true));
	}
	selectNone() {
	  this.selected.forEach((item) => item.removeAttribute("selected"));
	}
}

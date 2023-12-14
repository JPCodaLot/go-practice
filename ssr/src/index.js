import { SelectElement } from '../components/select';
import { CardElement } from '../components/card';

customElements.define("my-select", SelectElement);
customElements.define("my-card", CardElement);

window.addEventListener("DOMContentLoaded", () => {
  document.querySelector('#new').addEventListener("click", () => select.add());
  document.querySelector('#delete').addEventListener("click", () => select.removeSelected());
  document.querySelector('#select-all').addEventListener("click", () => select.selectAll());
  document.querySelector('#select-none').addEventListener("click", () => select.selectNone());
});

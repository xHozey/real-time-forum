const template = document.createElement("template");


template.innerHTML = `<ul>
</ul>`;

export class navBar extends HTMLElement {
  constructor() {
    super();
  }
  connectedCallback() {
    this.innerHTML = template.innerHTML;
  }
}

customElements.define("navBar-component", navBar);

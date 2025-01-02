const template = document.createElement('template')
template.innerHTML = `<div id="form">
        <a href="/register" class="href">Register</a>
        <h1>LOGIN</h1>
        <p>Username or email:</p>
        <input type="text" class="input-form">
        <p>Password:</p>
        <input type="password" class="input-form">
        <button class="btn">Submit</button>
    </div>`


export class login extends HTMLElement {
  constructor() {
    super();
  }
  connectedCallback() {
    this.innerHTML = template.innerHTML
  }

}

customElements.define('login-component', login)
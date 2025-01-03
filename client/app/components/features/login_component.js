const template = document.createElement("template");
template.innerHTML = `<div class="login">
        <a href="/register" class="href" data-link>Register</a>
        <h1>LOGIN</h1>
        <div class="container-eu">
        <p>Username or email:</p>
        <input type="text" class="input-login">
        </div>
        <div class="container-password">
        <p>Password:</p>
        <input type="password" class="input-login">
        </div>
        <button class="btn">Submit</button>
    </div>`;

export class login extends HTMLElement {
  constructor() {
    super();
  }
  connectedCallback() {
    this.innerHTML = template.innerHTML;
  }
}

customElements.define("login-component", login);

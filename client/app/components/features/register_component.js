const template = document.createElement('template')
template.innerHTML = `<div id="form">
        <a href="/login" class="href">Login</a>
        <h1>REGISTER</h1>
        <p>Nickname:</p>
        <input type="text" class="input-form">
        <p>Age:</p>
        <input type="number" class="input-form">
        <p>Gender:</p>
        <input type="radio" name="gender">
        <label for="gender">Male</label>
        <input type="radio" name="gender">
        <label for="gender">Female</label>
        <p>First Name:</p>
        <input type="text" class="input-form">
        <p>Last Name:</p>
        <input type="text" class="input-form">
        <p>E-mail:</p>
        <input type="email" class="input-form">
        <p>Password:</p>
        <input type="password" class="input-form">
        <button class="btn">Submit</button>
    </div>`


export class register extends HTMLElement {
  constructor() {
    super();
  }
  connectedCallback() {
    this.innerHTML = template.innerHTML
  }

}

customElements.define('register-component', register)
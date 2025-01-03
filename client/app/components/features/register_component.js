const template = document.createElement('template')
template.innerHTML = `<div class="register">
            <a href="/login" class="href" data-link>Login</a>
            <h1>REGISTER</h1>
            <div class="container-user">
                <p>Nickname:</p>    
                <input type="text" class="input-register">
            </div>
            <div class="container-age">
                <p>Age:</p>
                <input type="number" class="input-register">
            </div>
            <div class="container-gender">
                <p>Gender:</p>
                <input type="radio" name="gender">
                <label for="gender">Male</label>
                <input type="radio" name="gender">
                <label for="gender">Female</label>
            </div>
            <div class="container-fn">
                <p>First Name:</p>
                <input type="text" class="input-register">
            </div>
            <div class="container-ln">
                <p>Last Name:</p>
                <input type="text" class="input-register">
            </div>
            <div class="container-e">
                <p>E-mail:</p>
                <input type="email" class="input-register">
            </div>
            <div class="container-p">
                <p>Password:</p>
                <input type="password" class="input-register">
            </div>
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
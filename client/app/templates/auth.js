export const login = `
      <div class="container">
        <div class="login-box">
            <h1>Welcome back!</h1>
            <p>We're so excited to see you again!</p>
            <div class="loginForm">
                <label for="username">EMAIL OR NICKNAME</label>
                <input type="text" id="username" name="username">
                
                <label for="password">PASSWORD</label>
                <input type="password" id="password" name="password">

                <span class="error-message"></span>                
                <button id="login-btn">Login</button>
            </div>
            <p class="register-link">Need an account? <a href="/register" onclick="route()">Register</a></p>
        </div>
    </div>
`;

export const register = `
<div class="container">
        <div class="register-box">
            <h1>Create an Account</h1>
            <p>Join us today! It's quick and easy.</p>
            <div class="registerForm">
                <label for="nickname">NICKNAME</label>
                <input type="text" id="nickname" name="nickname">

                <label for="age">AGE</label>
                <input type="number" id="age" name="age">

                <label for="gender">GENDER</label>
                <select id="gender" name="gender">
                    <option value="" disabled selected>Select your gender</option>
                    <option value="male">Male</option>
                    <option value="female">Female</option>
                </select>

                <label for="firstName">FIRST NAME</label>
                <input type="text" id="firstName" name="firstName">

                <label for="lastName">LAST NAME</label>
                <input type="text" id="lastName" name="lastName">

                <label for="email">E-MAIL</label>
                <input type="email" id="email" name="email">

                <label for="password">PASSWORD</label>
                <input type="password" id="password" name="password">
                <span class="error-message"></span>
                <button id="register-btn" type="submit">Register</button>
            </div>
            <p class="login-link">Already have an account? <a href="/login" onclick="route()">Login</a></p>
        </div>
    </div>
`;

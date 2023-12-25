const loginForm = document.getElementById("login");
const registerForm = document.getElementById("register");
const toggleOptions = document.getElementById("btn");
const loginUsernameInput = document.getElementById("login-username");
const loginPasswordInput = document.getElementById("login-password");
const loginSubmitButton = loginForm.getElementsByClassName("submit-btn")[0];
const registerUsernameInput = document.getElementById("register-username");
const registerPasswordInput = document.getElementById("register-password");
const registerSubmitButton = registerForm.getElementsByClassName("submit-btn")[0];

registerUser();
loginUser();

function setRegister(){
    loginForm.style.left = "-400px";
    registerForm.style.left = "50px";
    toggleOptions.style.left = "110px";
}

function setLogin(){
    loginForm.style.left = "50px";
    registerForm.style.left = "450px";
    toggleOptions.style.left = "0";
}

function registerUser(){

    registerSubmitButton.addEventListener("click", (e) =>{
        e.preventDefault();

        const URL = "http://localhost:8080/register";
        const user = {
            id: 0,
            name: registerUsernameInput.value,
            password: registerPasswordInput.value,
        }

        fetch(URL, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(user)
        })
            .then(response => {
                if(!response.ok){
                    throw new Error("Network response was not ok");
                }
                setLogin();
                alert("Nutzer wurde erstellt. Melde dich nun an!")
                return response.json();
            })
            .then(data => {
                console.log("Response successful:", data);
            })
            .catch(error => {
                alert("Nutzer existiert bereits")
            })

    })
}

function loginUser(){
    loginSubmitButton.addEventListener("click", (e) => {
        e.preventDefault();

        const URL = "http://localhost:8080/login";
        const user = {
            id: 0,
            name: loginUsernameInput.value,
            password: loginPasswordInput.value,
        }

        fetch(URL, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(user)
        })
            .then(response => {
                if(!response.ok){
                    throw new Error("Network response was not ok");
                }
                window.location.href = "http://localhost:8080/todo-list/" + loginUsernameInput.value;
            })
            .catch(error => {
                alert("Nutzer konnte nicht angemeldet werden")
            })
    })
}
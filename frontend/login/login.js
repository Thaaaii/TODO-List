const loginForm = document.getElementById("login");
const registerForm = document.getElementById("register");
const toggleOptions = document.getElementById("btn");

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
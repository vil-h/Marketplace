

const form = document.querySelector(".register-form");

form.addEventListener("submit", function(event){
    event.preventDefault();




const email = document.querySelector("#email").value;
const password = document.querySelector("#password").value;


const data = {
    email,
    password
}

fetch("http://localhost:8080/register", {
    method: "POST",
    headers:{
        "Content-Type": "application/json",
    },

    body: JSON.stringify(data),
})

.then(response =>{
    window.location.href = "login.html";

})
});


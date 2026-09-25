const form = document.querySelector(".login-form");

form.addEventListener("submit", function(event){
    event.preventDefault();

    const email = document.querySelector("#email").value;
    const password = document.querySelector("#password").value;

    const data = {
        email,
        password
    }
    console.log(data);
    fetch("http://localhost:8080/login" ,{
        method: "POST",
        headers:{
            "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
    })

    .then(response =>{
        if (!response.ok){
            throw new Error("Неверный логин или пароль");
        }
        return response.json();
        
        

    })
    .then(data =>{  
        localStorage.setItem("token", data);
        window.location.href = "products.html";
    });
    

});
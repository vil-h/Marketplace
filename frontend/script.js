const searchInput = document.querySelector(".serch-input");



searchInput.addEventListener("input", function() {
    console.log(searchInput.value); //поиск товаров серч панель
});

fetch("http://localhost:8080/product")
    .then(response => {
        return response.json();
    })
    .then(products => {
        products.forEach(product => {
            const card = document.createElement("div");
            card.classList.add("product-card");

            const image = document.createElement("div");
            image.classList.add("product-image");
            card.append(image);

            const title = document.createElement("h2");
            title.textContent = product.name;   
            card.append(title);

            const price = document.createElement("p");
            price.textContent = product.price + " $";
            card.append(price);

            const button = document.createElement("button");
            button.textContent = "В корзину";
            card.append(button);

            const productsContainer = document.querySelector(".product");
            productsContainer.append(card);
        });;
    }); 



const token = localStorage.getItem("token")
fetch("http://localhost:8080/me",{
    headers: {
        "Authorization": `Bearer ${token}`
    }})

    .then(response => response.json())
    .then(data =>{
        console.log(data);
    const emailElement = document.querySelector("#user-email");
    emailElement.textContent = data.email;
    const roleElement = document.querySelector("#user-role");
    roleElement.textContent = data.role

    })
const loginButton = document.querySelector(".login-button");
const userPanel = document.querySelector(".user-panel");

if (token) {
    loginButton.style.display = "none";
    userPanel.style.display = "flex";
}

const logoutButton = document.querySelector("#logout-button");

logoutButton.addEventListener("click", function() {
    localStorage.removeItem("token");
    window.location.href = "login.html";
});

// Get the form element
let form = document.getElementById("user-form");

// Get the result element
let result = document.getElementById("result");

// Define what to do when the form is submitted
form.onsubmit = function (event) {
    // Prevent the default form submission behavior
    event.preventDefault();

    // Get the form data as an object
    let formData = {
        Name: form.name.value,
        Email: form.email.value,
        Company: form.company.value
    };

    // Convert the form data to JSON string
    let jsonData = JSON.stringify(formData);

    // Create a new XMLHttpRequest object
    let xhr = new XMLHttpRequest();

    // Set the request method and URL
    xhr.open("POST", "http://api:8080/api/users");

    // Set the request header for JSON content type
    xhr.setRequestHeader("Content-Type", "application/json");

    // Define what to do when the request is successful
    xhr.onload = function () {
        // Display the response text in the result element
        result.textContent = xhr.responseText;
    };

    // Send the request with the JSON data
    xhr.send(jsonData);
};

// Create a new XMLHttpRequest object
let xhr = new XMLHttpRequest();

// Set the request method and URL
xhr.open("GET", "http://localhost:8080/api/users");

// Define what to do when the request is successful
xhr.onload = function () {
    // Parse the response as JSON
    let data = JSON.parse(xhr.responseText);

    // Get the HTML element where we want to display the data
    let output = document.getElementById("users");

    // Loop through the data.data array and create a table row for each user
    let html = `
        <table class="table table-striped">
        <tr>
            <th>Name</th>
            <th>Email</th>
            <th>Company</th>
        </tr>`;
    for (let user of data.data) {
        html += `
        <tr>
        <td>${user.Name}</td>
        <td>${user.Email}</td>
        <td>${user.Company}</td>
        </tr>
        `;
    }
    html += "</table>";

    // Set the inner HTML of the output element to the generated table
    output.innerHTML = html;
};

// Send the request
xhr.send();

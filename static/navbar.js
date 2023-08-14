document.addEventListener("DOMContentLoaded", function () {
    const navbar = document.createElement("div");
    navbar.innerHTML = `
        <a href="/">Home</a>
        <a id="showAccount" href="/login"></a>
        <a id="loginLogout" href="/login">Log In</a>
    `;

    document.body.prepend(navbar);

    // Check if the session cookie exists
    const sessionCookie = getCookie("session");
    if (sessionCookie) {
        const loginLogout = document.getElementById("loginLogout");
        loginLogout.textContent = "Log Out";
        loginLogout.href = "/logout";
    }
    if (sessionCookie) {
        showAccount.textContent = "Account"
        showAccount.href = "/account"
    }

    // Function to get a cookie by name
    function getCookie(name) {
        const value = `; ${document.cookie}`;
        const parts = value.split(`; ${name}=`);
        if (parts.length === 2) return parts.pop().split(";").shift();
    }
});

document.getElementById("contactForm").addEventListener("submit", function(e) {
    e.preventDefault(); // prevent page reload

    const name = document.getElementById("name").value.trim();
    const email = document.getElementById("email").value.trim();
    const message = document.getElementById("message").value.trim();

    if (!name || !email || !message) {
        alert("Please fill in all fields.");
        return;
    }

    const subject = encodeURIComponent(`New message from ${name}`);
    const body = encodeURIComponent(`Name: ${name}\nEmail: ${email}\n\nMessage:\n${message}`);

    // Replace with your email
    const mailtoLink = `mailto:embadi43@gmail.com?subject=${subject}&body=${body}`;

    window.location.href = mailtoLink;
});

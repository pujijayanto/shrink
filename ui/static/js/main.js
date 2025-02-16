if (typeof originalURL !== 'undefined' && originalURL !== "" && originalURL !== "{{.OriginalURL}}") {
  document.getElementById('shortenForm').style.display = 'none';
  document.getElementById('redirecting').style.display = 'block';
  document.getElementById('redirecting').textContent = 'Redirecting to original URL...';
  setTimeout(() => {
      window.location.href = originalURL;
  }, 2000);
}

document.getElementById("shortenForm").addEventListener("submit", async (e) => {
  e.preventDefault();

  const formData = new FormData(e.target);
  const urlInput = document.getElementById("url");
  const submitButton = e.target.querySelector("button");

  try {
    const response = await fetch("/", {
      method: "POST",
      body: formData,
    });

    const data = await response.text();
    urlInput.value = data;

    // Change button to Copy
    submitButton.textContent = "Copy";
    submitButton.type = "button";

    // Add click handler for copy functionality
    submitButton.onclick = () => {
      urlInput.select();
      navigator.clipboard.writeText(urlInput.value);
      submitButton.textContent = "Copied!";
      setTimeout(() => {
        submitButton.textContent = "Copy";
      }, 2000);
    };

  } catch (error) {
    console.error("Error:", error);
    alert("error creating short url");
  }
});

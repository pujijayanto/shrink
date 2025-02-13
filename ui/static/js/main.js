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
    document.getElementById("result").innerHTML = "Error creating short URL";
  }
});

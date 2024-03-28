document.addEventListener("DOMContentLoaded", function() {
    document.getElementById("submitButton").addEventListener("click", async function(event) {
        event.preventDefault();
        var titleInput = document.getElementById("title");
        var bodyInput = document.getElementById("body");
        var title = titleInput.value;
        var body = bodyInput.value;
        var data = JSON.stringify({ "title": title, "body": body });

        try {
            const response = await fetch("/api/posts", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: data
            });
            
            if (!response.ok) {
                throw new Error("Failed to create post: " + response.statusText);
            }

            titleInput.value = "";
            bodyInput.value = "";
            
            Swal.fire({
                title: "Good job!",
                text: "Post added Successfully!",
                icon: "success",
                timer: 3000,
                timerProgressBar: true,
                onClose: () => {
                    window.location.href = "/";
                }
            });
            
        } catch (error) {
            alert(error.message);
        }
    });
});
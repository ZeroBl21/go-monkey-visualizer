document
	.getElementById("inputForm")
	.addEventListener("submit", async function (event) {
		event.preventDefault();

		// Obtener el texto del input
		const inputText = document.getElementById("inputText").value;

		const response = await fetch("http://localhost:5173/api/lexer", {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify({ input: inputText }),
		});

		if (response instanceof Error) {
			document.getElementById("outputText").value = "Error: " + error.message;
		}

		if (!response.ok) {
			document.getElementById("outputText").value =
				`HTTP Error ${response.status}: Unable to process input.`;
		}

		const data = await response.json().catch((err) => err);

		if (data instanceof Error) {
			console.error("Response parsing error:", data.message);
			return { isSuccess: false, error: data.message };
		}

		document.getElementById("outputText").value = JSON.stringify(
			data.result,
			null,
			2,
		);
	});

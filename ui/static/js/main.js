document
	.getElementById("inputForm")
	.addEventListener("submit", async function (event) {
		event.preventDefault();

		const inputText = document.getElementById("inputText").value.trim();
		const processType = document.getElementById("processType").value;

		if (inputText === "") {
			document.getElementById("outputText").value =
				"Error: El campo de entrada no puede estar vacío.";
			return;
		}

		const url =
			processType === "flex"
				? "http://localhost:5173/api/flex"
				: "http://localhost:5173/api/lexer";

		const response = await fetch(url, {
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

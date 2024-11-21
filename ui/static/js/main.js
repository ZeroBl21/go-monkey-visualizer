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

		let url;
		if (processType === "monkey") {
			url = "http://localhost:5173/api/lexer";
		} else if (processType === "pratt") {
			url = "http://localhost:5173/api/pratt";
		} else if (processType === "evaluator") {
			url = "http://localhost:5173/api/evaluator";
		} else if (processType === "bytecode") {
			url = "http://localhost:5173/api/compiler";
		} else {
			url = "http://localhost:5173/api/lexer";
		}

		const response = await fetch(url, {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify({ input: inputText }),
		}).catch((err) => err);

		if (response instanceof Error) {
			document.getElementById("outputText").value = "Error: " + error.message;
			return;
		}

		if (!response.ok) {
			document.getElementById("outputText").value =
				`HTTP Error ${response.status}: Unable to process input.`;
			return;
		}

		const data = await response.json().catch((err) => err);

		if (data instanceof Error) {
			console.error("Response parsing error:", data.message);
			document.getElementById("outputText").value = "Error: " + error.message;
			return;
		}

		document.getElementById("outputText").value = JSON.stringify(
			data.result,
			null,
			2,
		);
	});

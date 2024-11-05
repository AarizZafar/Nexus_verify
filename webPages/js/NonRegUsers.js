async function fetchRecords() {
    try {
        const response = await fetch("/GetUnregBM");
        const records = await response.json();

        const container = document.getElementById("records-container");
        container.innerHTML = "";

        records.forEach(record => {
            const block = document.createElement("div");
            block.className = "block";
            block.innerHTML = `
                <div class="title">TestNet: ${record.testNet || 'N/A'}</div>
                <div>SSID: ${record.sysBiometric?.ssid || 'N/A'}</div>
                <div>MAC: ${record.sysBiometric?.mac || 'N/A'}</div>
                <div>Serial: ${record.sysBiometric?.systemserialnumber || 'N/A'}</div>
                <div>UUID: ${record.sysBiometric?.uuid || 'N/A'}</div>
            `;
            block.addEventListener("click", () => sendClickedRecord(record, block));
            container.appendChild(block);
        });
    } catch (error) {
        console.error("Error fetching records: ", error);
    }
}

async function sendClickedRecord(record, block) {
    try {
        await fetch("/SetClickedRecord", {
            method : "POST",
            header: {"Content-Type" : "application/json"},
            body : JSON.stringify(record)
        });
        console.log("Record sent to Go server: ", record);
        block.remove();
    } catch(error) {
        console.error("Error sending record : ", error);
    }
}

fetchRecords()
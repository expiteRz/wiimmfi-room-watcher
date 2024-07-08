const queryParams = new URLSearchParams(window.location.search);

const tab = +(queryParams.get('tab') || 0);

function snackbarGenerate(status, text) {
    let elm = document.createElement("div");
    elm.classList.add("toast");
    elm.classList.add(status);
    elm.textContent = text;
    document.body.appendChild(elm);
    setTimeout(() => {
        const exist_toast = document.body.querySelector(".toast");
        exist_toast.classList.add("show");
    }, 100);
    setTimeout(() => {
        const exist_toast = document.body.querySelector(".toast");
        exist_toast.classList.remove("show")
        setTimeout(() => {
            exist_toast.remove();
        }, 1000);
    }, 3000);
}

document.querySelectorAll(".nav>.button").forEach(item => {
    if (!item.href.includes(`?tab=${tab}`)) return;
    item.classList.add("active");
});

window.addEventListener("click", async ev => {
    const t = ev.target;

    if (t?.id === "save_settings") {
        const interval = document.querySelector("#interval");
        const interval_value = parseInt(interval.value);
        const interval_min = parseInt(interval.min);
        interval.value = interval_value < interval_min ? interval.min : interval.value;
        const download = await fetch("/api/setting/save", {
            method: "POST",
            body: JSON.stringify({
                pid: document.querySelector("#pid").value,
                server_ip: document.querySelector("#serverip").value,
                interval: document.querySelector("#interval").value,
                server_port: document.querySelector("#serverport").value
                // server_port: "omg"
            })
        });
        const res = await download.json();
        if (res.status !== "success") {
            snackbarGenerate("error", "Saving setting failed. Please check error on console");
            return;
        }
        if (res.need_restart) snackbarGenerate("success", "Setting saved. You need to restart the server to apply your update");
        else snackbarGenerate("success", "Setting saved successfully.");
    }

    if (t?.classList.contains("copy")) {
        if (t.attributes.to?.value === undefined || t.attributes.to?.value === null) {
            snackbarGenerate("error", "Copy failed. URL is invalid.");
            return;
        }
        await navigator.clipboard.writeText(t.attributes.to?.value);
        snackbarGenerate("success", "URL copied successfully.");
    }

    if (t?.classList.contains("open_button")) return open_folder(t);
});

async function open_folder(el) {
    const folderName = decodeURI(el.attributes.to?.value || "");

    const download = await fetch(`/api/overlays/open/${folderName}`);
}
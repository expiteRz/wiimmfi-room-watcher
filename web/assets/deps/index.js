const queryParams = new URLSearchParams(window.location.search);

const tab = +(queryParams.get('tab') || 0);

document.querySelectorAll(".nav>.button").forEach(item => {
    if (!item.href.includes(`?tab=${tab}`)) return;
    item.classList.add("active");
});

window.addEventListener("click", async ev => {
    const t = ev.target;

    if (t?.id === "save_settings") {
        const download = await fetch("/api/setting/save", {
            method: "POST",
            body: JSON.stringify({
                pid: document.querySelector("#pid").value,
                server_ip: document.querySelector("#serverip").value,
                interval: document.querySelector("#interval").value
            })
        });
        const res = await download.text();
        console.log(res);
    }
});
var target = document.querySelector('.chat');
target.scrollTo(0, target.scrollHeight);
function addMessage() {
    var message = document.getElementById("input__message").value;
    message = message.trim();
    var target = document.querySelector('.chat');
    target.scrollTo(0, target.scrollHeight);
    if (message != "") {
        let now = new Date();
        var month = now.getMonth;
        target.innerHTML += `
        <div class="examp_message you">
        <div class="top">
            <div class="name">
                <p class="ico"></p>
                <p>Вы</p>
            </div>
            <div class="date"><span>${now.getDay()+1 < 10 ? 0 : ""}${now.getDay() + 1}.${now.getMonth()+1 < 10 ? 0 : ""}${now.getMonth() + 1} ${now.getHours()}:${now.getMinutes()}</span></div>
        </div>
        <p class="message">${message}
        </p>
    </div>
    `
    }
}
let answerShowing = true;
const $consoleText = $("#console-text");

$(".operator-number").on("click", function () {
    let $clickedButton = $.trim($(this).text());
    $(this).fadeOut(100).fadeIn(100);

    if (answerShowing && $clickedButton !== "Back" && $clickedButton !== '=') {
        if (!['+', '-', '/', '*'].includes($clickedButton) || $consoleText.text() === "Waiting For Input") {
            $consoleText.text("");
        }
        answerShowing = false;
    }
    if ($clickedButton === '=') {
        if (!answerShowing) {
            consoleEval();
            answerShowing = true;
        }
        return;
    } else if ($clickedButton === "Clear") {
        $consoleText.text("");
        answerShowing = true;
        return;
    } else if ($clickedButton === "Back") {
        if ($consoleText.text() !== "Waiting For Input") {
            $consoleText.text($consoleText.text().slice(0, -1));
            answerShowing = false;
        }
        return;
    }
    $consoleText.text($consoleText.text() + $clickedButton);
});


function consoleEval() {
    $.ajax({
        type: "POST",
        url: "/calculate",
        data: $consoleText.text(),
        contentType: "text/plain",
        success: (response) => {
            $consoleText.text(response);
        },
        error: (xhr, status, error) => {
            $consoleText.text("ERROR");
        },
    });
}
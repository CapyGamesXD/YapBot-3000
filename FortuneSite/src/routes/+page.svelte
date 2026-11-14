<script>

let reply = $state('')
async function generate() {
    const response = await fetch('fortune.capydesigns.com')
    reply = await response.json()
}
</script>

<div class="middle">
<h1 style="margin-top: 10px;">YapBot 3000</h1>
<h2 class="fancy">Your badly generated fortune awaits</h2>
<p style="margin-top: 30px;">Try it now!</p>
<p>Click to generate your weird fortune...</p>
<button style="margin-top: 20px;" class="genButton" onclick={generate}>Generate</button>
<p>{reply}</p>
</div>

<style>
    :global(body) {
        background-color: rgb(31, 31, 31);
    }
    h1 {
        font-size: 30px;
        font-family: "Koulen", sans-serif;
        padding: 0;
        margin: 0;
        color: white;
    }
    h2 {
        font-size: 20px;
        font-family: "Jost", sans-serif;
        font-weight: 300;
        padding: 0;
        margin: 0;
        color: white;
    }
    p {
        font-family: "Jost", sans-serif;
        color: white;
        padding: 0;
        margin: 0;
    }
    .middle {
        display: flex;
        align-items: center;
        justify-content: center;
        flex-direction: column;
    }
    .genButton {
        border: none;
        outline: none;
        border-radius: 10px;
        width: 200px;
        height: 80px;
        font-size: 20px;
        background-color: rgb(150, 134, 216);
        box-shadow: inset 0px -6px 0px rgb(126, 112, 185);
        padding-bottom: 6px;
        transition: 0.15s;
    }
    .genButton:hover {
        background-color: rgb(150, 134, 216);
        transform: scale(1.05);
    }
    .genButton:active {
        box-shadow: none;
        padding-bottom: 0;
    }
</style>
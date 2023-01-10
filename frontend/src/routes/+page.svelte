<script>
    let url = '';
    let shortenedUrl = '';

    const handleSubmission = () => {
        console.log('Submitted');
        url = '';
        shortenedUrl = 'https://www.google.com';
    }

</script>

<div class="container">
    <h1>Shorten Any Link</h1>
    <label for="url">URL</label>
    <input type="text" id="url" name="url" placeholder="Enter URL" bind:value={url} />
    <button on:click={handleSubmission}>Shorten</button>
    {#if shortenedUrl}
        <p>Shortened URL: <a href={shortenedUrl}>{shortenedUrl}</a></p>
    {/if}
</div>

<style>
.container {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 100vh;
    width: 100vw;
    background-color: #f5f5f5;
}

.container h1 {
    font-size: 2rem;
    font-weight: 500;
    margin-bottom: 1rem;
}

.container label{
    font-size: 1.5rem;
    font-weight: 500;
    margin-bottom: 0.5rem;
}

.container input{
    width: 50%;
    height: 2rem;
    border: 1px solid #ccc;
    border-radius: 5px;
    padding: 0.5rem;
    font-size: 1rem;
    margin-bottom: 1rem;
}

.container button{
    width: 50%;
    height: 2rem;
    border: 1px solid #ccc;
    border-radius: 5px;
    padding: 0.5rem;
    font-size: 1rem;
    background-color: #fff;
    cursor: pointer;
}

.container p{
    font-size: 1.5rem;
    font-weight: 500;
    margin-top: 1rem;
}

</style>



/** @type {import('./$types').PageLoad} */
export async function load({ params, fetch }) {
	const shortUrl = params.url;
	let longUrl = 'invalid url';

	try {
        const res = await fetch(`http://localhost:3000/${shortUrl}`);
        const data = await res.json();
        longUrl = data.url;
    } catch (err) {
        console.log(err);
    }


	return {
		props: {
			url: longUrl
		}
	};
}

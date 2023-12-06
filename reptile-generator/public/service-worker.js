async function addResourcesToCache(resources) {
	const cache = await caches.open("v1");
	await cache.addAll(resources);
}

self.addEventListener("install", (event) => {
	event.waitUntil(
		addResourcesToCache([
			"https://upload.wikimedia.org/wikipedia/commons/thumb/1/18/Bartagame_fcm.jpg/640px-Bartagame_fcm.jpg",
			"https://upload.wikimedia.org/wikipedia/commons/4/4d/Ball_python_lucy.JPG",
		])
	);
});

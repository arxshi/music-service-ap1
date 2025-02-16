document.addEventListener("DOMContentLoaded", function () {
    const audioPlayer = document.querySelector("audio");
    const songTitle = document.querySelector(".song");
    const songArtist = document.querySelector(".artist");
    const songImage = document.querySelector(".footer img");
    const searchInput = document.getElementById("searchInput");
    const searchButton = document.getElementById("searchButton");
    const apiUrl = "/tracks"; // Обновленный эндпоинт

    let playlist = [];

    // Функция загрузки списка песен с сервера
    function loadMusic() {
        fetch(apiUrl)
            .then(response => response.json())
            .then(data => {
                playlist = data;
                updatePlaylistUI();
            })
            .catch(error => console.error("Ошибка загрузки музыки:", error));
    }

    function updatePlaylistUI() {
        const musicList = document.getElementById("music-list");
        musicList.innerHTML = "";
        playlist.forEach((song, index) => {
            const item = document.createElement("li");
            item.textContent = `${song.artist} - ${song.title}`;
            item.addEventListener("click", () => playSong(index));
            musicList.appendChild(item);
        });
    }

    // Функция добавления новой песни
    document.getElementById("add-music-form").addEventListener("submit", function (event) {
        event.preventDefault();
        const artist = document.getElementById("artist").value;
        const title = document.getElementById("title").value;

        fetch(apiUrl, {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ artist, title })
        })
            .then(response => response.json())
            .then(() => {
                loadMusic(); // Обновляем список
            })
            .catch(error => console.error("Ошибка добавления музыки:", error));
    });

    function playSong(index) {
        const song = playlist[index];
        songTitle.textContent = song.title;
        songArtist.textContent = song.artist;
        audioPlayer.src = song.url;
        audioPlayer.play();
    }

    loadMusic();
});

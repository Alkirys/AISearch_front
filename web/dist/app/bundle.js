// - Рабинович, какими вы видите перспективы этого проекта?
// - Таки судя по значительной части его разработчиков - весьма радужными...

// let canvas = document.querySelector('canvas');
// let context = canvas.getContext('2d');

const USER_IDS_EVENT = 'USER_IDS_EVENT';
const LIST_LOADING_STATE_CHANGE = 'LIST_LOADING_STATE_CHANGE';
const NO_FACES = 'NO_FACES';
const FETCH_ERROR = 'FETCH_ERROR';

class EventBus {
    constructor() {
        this.events = {};
    }

    on(evt, callback) {
        (this.events[evt] || (this.events[evt] = [])).push(callback);
    }

    off(evt, callback) {
        if (!this.events[evt]) {
            return;
        }
        this.events[evt] = this.events[evt].filter((func) => func !== callback);
    }

    emit(evt, arg) {
        (this.events[evt] || []).slice().forEach((lsn) => lsn(arg));
    }
}

const eventBus = new EventBus();
let VK_access_token = '';

fetch("config.json").then((res) => {
    return res.json().then((config) => {
        VK_access_token = config.VK_access_token;
    });
});

class App {
    constructor() {
        this.place = document.getElementById('mainContainer');
        this.header = new Header();
        this.photoContainer = new PhotoContainer();
        this.usersList = new UsersList();
    }

    getContainerMaxHeight() {
        const windowHeight = window.innerHeight
            || document.documentElement.clientHeight
            || document.body.clientHeight;

        const containerMaxHeight = windowHeight - 290; // тут чисто хардкод под стили для css-класса mainContainer
        return containerMaxHeight;
    }

    render() {
        const photoMaxHeight = this.getContainerMaxHeight();
        this.place.innerHTML = `
            <header id="header" class="header"></header>
            <div id="photoContainer" class="photoContainer" style="height: ${photoMaxHeight}px"></div>
            <div class="listWrapper" id="listWrapper"></div>
        `;

        this.header.setPlace(document.getElementById('header'));
        this.photoContainer.setPlace(document.getElementById('photoContainer'));
        this.usersList.setPlace(document.getElementById('listWrapper'));

        this.header.render();
        this.photoContainer.render();
        this.usersList.render();
    }
}

class UsersList {
    constructor() {
        this.place = '';
        this.isLoading = false;
        this.error = false;
    }

    setPlace(place) {
        this.place = place;
    }

    vkGetUsersCallback(usersData) {
        if (!usersData.response) {
            eventBus.emit(FETCH_ERROR);
        } else {
            this.usersList = usersData.response.map((userData) => {
                return {id: userData.id, avatarUrl: userData.photo_100, name: userData.first_name + ' ' + userData.last_name};
            });
            eventBus.emit(LIST_LOADING_STATE_CHANGE);
        }
        const script = document.getElementById('vkApiScript');
        if (script) {
            script.outerHTML = '';
        }
    }

    render() {
        // TODO порефакторить условия
        if (!this.usersList && !this.isLoading && !this.error) {
            this.place.innerHTML = `
                <div id="notFaceFound" class="notFaceFound">
                    <p class="notFaceFound__message">Загрузите фото...</p>
                </div>
            `;
            this._addListeners();
            this._subscribeEvents();
            return;
        }

        if (this.error) {
            this.place.innerHTML = `
                <div id="notFaceFound" class="notFaceFound">
                    <p class="notFaceFound__message">Кажется, что-то пошло не так... Попробуйте еще раз</p>
                </div>
            `;
            this._addListeners();
            return;
        }

        if (this.isLoading) {
            this.place.innerHTML = `
                <div id="notFaceFound" class="notFaceFound">
                    <div class="spinner">
                        <div></div><div></div><div></div><div></div><div></div><div></div><div></div><div></div>
                    </div>
                </div>
            `;
            this._addListeners();
            return;
        }

        if (!this.usersList && this.usersList.length === 0 && !this.isLoading) {
            this.place.innerHTML = `
                <div id="notFaceFound" class="notFaceFound">
                    <p class="notFaceFound__message">Не удалось разобрать лицо на фото :(</p>
                </div>
            `;
            this._addListeners();
            return;
        }

        const usersHtml = this.usersList.reduce((htmlStr, userData) => {
            return htmlStr + `
                <li class="list__elem">
                    <a href="https://vk.com/id${userData.id}" target="_blank" class="list__link">
                        <img src="${userData.avatarUrl}" class="list__avatar" alt="avatar"/>
                        <p class="list__name">${userData.name}</p>
                    </a>
                </li>
            `;
        }, '');

        this.place.innerHTML = `
            <ul id="list" class="list">
                <div class="list__container">
                    ${usersHtml}
                </div>
            </ul>
        `;

        this._addListeners();
    }

    _rerenderLoading = () => {
        this.error = false;
        this.isLoading = !this.isLoading;
        this._removeListeners();
        this.render();
    }

    _rerenderError = () => {
        this.error = true;
        this.isLoading = false;
        this._removeListeners();
        this.render();
    }

    _addListeners() {
    }

    _removeListeners() {
    }

    _subscribeEvents() {
        eventBus.on(USER_IDS_EVENT, this._fetchVKInfo);
        eventBus.on(LIST_LOADING_STATE_CHANGE, this._rerenderLoading);
        eventBus.on(NO_FACES, this.vkGetUsersCallback);
        eventBus.on(FETCH_ERROR, this._rerenderError);
    }

    _unsubscribeEvents() {
        eventBus.off(USER_IDS_EVENT, this._fetchVKInfo);
        eventBus.off(LIST_LOADING_STATE_CHANGE, this._rerenderLoading);
        eventBus.off(NO_FACES, this.vkGetUsersCallback);
        eventBus.off(FETCH_ERROR, this._rerenderError);
    }

    _fetchVKInfo = (data) => {
        const script = document.createElement('script');
        script.id = 'vkApiScript';
        script.src = `https://api.vk.com/method/users.get?user_ids=${data}&fields=photo_100&access_token=${VK_access_token}&v=5.131&callback=app.usersList.vkGetUsersCallback`;
        document.getElementsByTagName("head")[0].appendChild(script);
    }

    hide() {
        this._unsubscribeEvents();
        this._removeListeners();
        this.place.innerHTML = '';
    }
}

class Header {
    constructor() {
        this.place = '';
    }

    setPlace(place) {
        this.place = place;
    }

    render() {
        this.place.innerHTML = `
                <div class="header__buttonGroup">
                    <p id="mainButton" class="header__button header__activeButton">AISearch</p>
                    <p id="aboutButton" class="header__button header__passiveButton">about us</p>
                </div>
                <div class="header__buttonGroup">
                    <a class="header__iconGroup header__passiveButton" href="https://t.me/InternalServerAI" target="_blank">
                        <p id="contactButton" class="header__button">contact us</p>
                        <svg width="32" height="32" viewBox="0 0 32 32" fill="none">
                            <path d="M28.4012 3.83123C28.0895 3.8449 27.7861 3.93166 27.5125 4.04123C27.2417 4.15007 25.6825 4.81556 23.38 5.79998C21.0775 6.7844 18.0844 8.06633 15.1175 9.33748C9.18357 11.8798 3.35249 14.3812 3.35249 14.3812L3.39499 14.365C3.39499 14.365 3.04305 14.4826 2.68624 14.7312C2.50783 14.8555 2.31589 15.016 2.15874 15.25C2.00158 15.4839 1.88669 15.8155 1.92999 16.1662C2.08206 17.3981 3.35874 17.7475 3.35874 17.7475L3.36374 17.75L9.06124 19.7C9.2065 20.1846 10.7893 25.4674 11.1375 26.5862C11.3298 27.2048 11.5098 27.5616 11.6975 27.805C11.7914 27.9268 11.8899 28.021 11.9987 28.09C12.042 28.1174 12.0875 28.1379 12.1325 28.1562C12.1334 28.1567 12.134 28.1558 12.135 28.1562C12.1405 28.1587 12.1457 28.1589 12.1512 28.1612L12.1362 28.1575C12.1467 28.1617 12.157 28.1688 12.1675 28.1725C12.1878 28.1796 12.2002 28.1794 12.2262 28.185C12.8787 28.4128 13.4187 27.99 13.4187 27.99L13.4412 27.9725L16.9412 24.7262L22.6212 29.155L22.6925 29.1875C23.6867 29.6289 24.5858 29.3826 25.0837 28.9775C25.5817 28.5723 25.7775 28.05 25.7775 28.05L25.7987 27.995L29.9725 6.24123C30.0793 5.75455 30.0945 5.33387 29.9937 4.95373C29.893 4.57359 29.6483 4.24314 29.3412 4.05998C29.0342 3.87682 28.7129 3.81755 28.4012 3.83123ZM28.435 5.13248C28.5616 5.12662 28.6554 5.14109 28.685 5.15873C28.7146 5.17637 28.7278 5.17405 28.7562 5.28123C28.7846 5.38841 28.8 5.61416 28.7225 5.96748L28.72 5.97498L24.57 27.6025C24.5602 27.6243 24.4737 27.8243 24.2762 27.985C24.0748 28.1489 23.8508 28.2768 23.2562 28.0275L17.045 23.1837L16.87 23.0462L16.8662 23.05L15.0087 21.6562L25.4475 9.37498C25.5277 9.28086 25.5789 9.16551 25.595 9.04289C25.611 8.92027 25.5912 8.79563 25.5379 8.68403C25.4846 8.57244 25.4002 8.47867 25.2947 8.41406C25.1893 8.34945 25.0674 8.31677 24.9437 8.31998C24.8229 8.32312 24.7055 8.36039 24.605 8.42748L9.49999 18.4975L3.79374 16.5437C3.79374 16.5437 3.22711 16.2284 3.19999 16.0087C3.19848 15.9966 3.19179 16.0076 3.22124 15.9637C3.25068 15.9199 3.32469 15.8459 3.41749 15.7812C3.60307 15.6519 3.81499 15.5737 3.81499 15.5737L3.83624 15.5662L3.85749 15.5575C3.85749 15.5575 9.68888 13.0559 15.6225 10.5137C18.5893 9.24264 21.5818 7.96166 23.8837 6.97748C26.1851 5.99355 27.8479 5.28533 27.9887 5.22873C28.149 5.16454 28.3084 5.13834 28.435 5.13248ZM21.5125 12.0275L13.5962 21.3412L13.5925 21.345C13.5801 21.3599 13.5684 21.3753 13.5575 21.3912C13.5449 21.4086 13.5332 21.4265 13.5225 21.445C13.478 21.5205 13.4495 21.6043 13.4387 21.6912C13.4387 21.6929 13.4387 21.6946 13.4387 21.6962L12.4075 26.3337C12.3903 26.2836 12.3783 26.2653 12.36 26.2062V26.205C12.0324 25.1525 10.5391 20.1703 10.3325 19.4812L21.5125 12.0275ZM14.49 22.8675L15.9125 23.935L13.8225 25.8725L14.49 22.8675Z" fill="white"/>
                        </svg>
                    </a>
                </div>
        `;
    }

    hide() {
        this.place.innerHTML = '';
    }
}

class PhotoContainer {
    constructor() {
        this.place = '';
        this.photoPlaceholder = new PhotoPlaceholder();
        this.fileInput = new FileInput();
    }

    setPlace(place) {
        this.place = place;
    }

    render() {
        const photoMaxWidth = this.getContainerMaxHeight(); // считаем, что фото - квадрат, тогда макс ширина равна макс длинне
        this.place.innerHTML = `
                <div id="photo" class="photo" style="width: ${photoMaxWidth}px"></div>
                <div id="input" class="input"></div>
        `;
        this.photoPlaceholder.setPlace(document.getElementById('photo'));
        this.fileInput.setPlace(document.getElementById('input'));
        this.photoPlaceholder.render();
        this.fileInput.render();
    }

    // TODO метод часто повторяется
    getContainerMaxHeight() {
        const windowHeight = window.innerHeight
            || document.documentElement.clientHeight
            || document.body.clientHeight;

        const containerMaxHeight = windowHeight - 290; // тут чисто хардкод под стили для css-класса mainContainer
        return containerMaxHeight;
    }

    hide() {
        this.place.innerHTML = '';
    }
}

class FileInput {
    constructor() {
        this.place = '';
        this.fileInput = '';
        this.photo = new Photo();
    }

    setPlace(place) {
        this.place = place;
    }

    _fileInputChangeHandler = (evt) => {
        const file = evt.target.files[0];
        if (!file.type.match('image.*')) {
            alert('wrong file type');
        }
        const fileReader = new FileReader();
        fileReader.onload = (function(file) {
            return (e) => {
                this.photo.setPlace(document.getElementById('photo'));
                this.photo.setFileData(file.fileName, e.target.result);
                this.photo.render();
            };
        }).bind(this)(file);
        fileReader.readAsDataURL(file);
    }

    render() {
        this.place.innerHTML = `
            <input id="inputFile" type="file" class="input__inputElement">
            <label for="inputFile" class="input__label">
                <div class="input__animatedButton">
                <svg xmlns="http://www.w3.org/2000/svg" width="45" height="45" viewBox="5.7 5.7 45.6 45.6" fill="none" class="input__icon">
                    <path fill-rule="evenodd" clip-rule="evenodd" d="M28.5 17.8125C22.5975 17.8125 17.8125 22.5974 17.8125 28.5C17.8125 34.4025 22.5975 39.1875 28.5 39.1875C34.4026 39.1875 39.1875 34.4025 39.1875 28.5C39.1875 22.5974 34.4026 17.8125 28.5 17.8125ZM22.0875 28.5C22.0875 24.9584 24.9585 22.0875 28.5 22.0875C32.0415 22.0875 34.9125 24.9584 34.9125 28.5C34.9125 32.0415 32.0415 34.9125 28.5 34.9125C24.9585 34.9125 22.0875 32.0415 22.0875 28.5Z" fill="white"/>                     <path fill-rule="evenodd" clip-rule="evenodd" d="M5.70001 23.9399C5.70001 17.5554 5.70001 14.3631 6.94254 11.9245C8.03549 9.77941 9.77947 8.03543 11.9245 6.94248C14.3631 5.69995 17.5554 5.69995 23.94 5.69995H33.06C39.4446 5.69995 42.6369 5.69995 45.0755 6.94248C47.2206 8.03543 48.9645 9.77941 50.0575 11.9245C51.3 14.3631 51.3 17.5554 51.3 23.94V33.06C51.3 39.4445 51.3 42.6368 50.0575 45.0754C48.9645 47.2205 47.2206 48.9645 45.0755 50.0574C42.6369 51.2999 39.4446 51.2999 33.06 51.2999H23.94C17.5554 51.2999 14.3631 51.2999 11.9245 50.0574C9.77947 48.9645 8.03549 47.2205 6.94254 45.0754C5.70001 42.6368 5.70001 39.4445 5.70001 33.06V23.9399ZM23.94 9.97495H33.06C36.3229 9.97495 38.5129 9.97828 40.1994 10.1161C41.836 10.2498 42.6185 10.4885 43.1347 10.7515C44.4753 11.4346 45.5653 12.5246 46.2484 13.8653C46.5114 14.3815 46.7502 15.164 46.8839 16.8006C47.0217 18.4871 47.025 20.6771 47.025 23.94V33.06C47.025 36.3228 47.0217 38.5128 46.8839 40.1993C46.7502 41.8359 46.5114 42.6184 46.2484 43.1346C45.5653 44.4753 44.4753 45.5653 43.1347 46.2484C42.6185 46.5114 41.836 46.7501 40.1994 46.8838C38.5129 47.0216 36.3229 47.0249 33.06 47.0249H23.94C20.6772 47.0249 18.4871 47.0216 16.8006 46.8838C15.164 46.7501 14.3815 46.5114 13.8653 46.2484C12.5247 45.5653 11.4347 44.4753 10.7516 43.1346C10.4886 42.6184 10.2498 41.8359 10.1161 40.1993C9.97834 38.5128 9.97501 36.3228 9.97501 33.06V23.9399C9.97501 20.6771 9.97834 18.4871 10.1161 16.8006C10.2498 15.164 10.4886 14.3815 10.7516 13.8653C11.4347 12.5246 12.5247 11.4346 13.8653 10.7515C14.3815 10.4885 15.164 10.2498 16.8006 10.1161C18.4871 9.97828 20.6772 9.97495 23.94 9.97495Z" fill="white"/>
                </svg>
                </div>
                <span class="input__button input__animatedButton">Выбрать</span>
            </label>
            <p class="input__description">Чтобы увидеть список пользователей, загрузите фото...</p>
        `;

        this.fileInput = document.getElementById('inputFile');
        this.fileInput.addEventListener('change', this._fileInputChangeHandler);
    }

    hide() {
        this.fileInput.removeEventListener('change', this._fileInputChangeHandler);
        this.place.innerHTML = '';
    }
}

class Photo {
    constructor() {
        this.place = '';
        this.facesData = [];
    }

    setPlace(place) {
        this.place = place;
    }

    setFileData(fileName, base64) {
        this.fileName = fileName;
        this.base64 = base64;
    }

    render() {
        if (!this.base64) {
            console.log('no file src');
            return;
        }

        this.place.innerHTML = `
            <canvas class="photo__canvas"></canvas>
        `;

        this.canvas = document.querySelector('canvas');
        const context = this.canvas.getContext('2d');
        const image = new Image();
        image.src = this.base64;
        image.onload = () => {
            const imgSize = this.getImageSize(image.width, image.height);
            this.canvas.height = imgSize.height;
            this.canvas.width = imgSize.width;
            context.drawImage(image,0, 0, imgSize.width, imgSize.height);
            eventBus.emit(LIST_LOADING_STATE_CHANGE);
            this.findFacesRequest();
        };
    }

    findFacesRequest() {
        if (!this.canvas) {
            console.log('no image loaded');
            return;
        }
        let b64Text = this.canvas.toDataURL('image/jpeg');
        b64Text = b64Text.replace('data:image/jpeg;base64,','');

        fetch('http://127.0.0.1/api/v1/find', {
            method: "POST",
            body: JSON.stringify({img: b64Text})
        }).then((res) => {
            res.json().then((data) => {
                this.facesData = data.imgs;
                this.detectFacesRequest();
            });
        }).catch(() => eventBus.emit(FETCH_ERROR));
    }

    detectFacesRequest() {
        if (!this.facesData) {
            console.log('no find request');
            return;
        }
        if (this.facesData.length === 0) {
            eventBus.emit(NO_FACES, {response: []});
            console.log('no facesFound');
            return;
        }
        fetch('http://127.0.0.1/api/v1/detect', {
            method: "POST",
            body: JSON.stringify({img: this.facesData[0]}) // TODO пока для одного лица
        }).then((res) => {
            res.json().then((data) => {
                this.userIds = this.validateUserUrls(data.users);
                eventBus.emit(USER_IDS_EVENT, this.userIds);
            });
        }).catch(() => eventBus.emit(FETCH_ERROR));
    }

    validateUserUrls(users) {
        const idsArray = users.map((userUrl) => {
            return userUrl.split('id')[1];
        });
        const idsStr = idsArray.reduce((resStr, id) => resStr + id + ',', '');
        return idsStr.slice(0, idsStr.length - 1);
    }

    getImageSize(imageWidth, imageHeight) {
        const windowHeight = window.innerHeight
            || document.documentElement.clientHeight
            || document.body.clientHeight;

        const imageMaxHeight = windowHeight - 290; // тут чисто хардкод под стили для css-класса mainContainer
        const imageAspectRatio = imageWidth / imageHeight;
        // тк должна быть константрная ширина, считаем, что картинка квадратная и берем макс ширину = макс высоте
        let realImageWidth = 0;
        let realImageHeight = 0;
        if (imageWidth > imageHeight) {
            realImageWidth = imageMaxHeight;
            realImageHeight = realImageWidth * imageHeight / imageWidth;
        } else {
            realImageHeight = imageMaxHeight;
            realImageWidth = realImageHeight * imageWidth / imageHeight;
        }
        return {width: realImageWidth, height: realImageHeight};
    }

    hide() {
        this.place.innerHTML = '';
    }
}

class PhotoPlaceholder {
    constructor() {
        this.place = '';
    }

    setPlace(place) {
        this.place = place;
    }

    render() {
        const svgSize = {width: 527, height: 519}; // это примерные размеры картинки в svg без масштабирования
        const svgMaxWidth = this.getSvgMaxHeight();
        const svgMaxHeight = svgMaxWidth * svgSize.height / svgSize.width;
        this.place.innerHTML = `
            <svg class="photo__placeholder" width="${svgMaxWidth}" height="${svgMaxHeight}" viewBox="75 5 550 540" fill="" fill-opacity="0.15" stroke="white">
                <path d="m166.62 437.83 20.176 3.5391-6.0547 49.641h0.003907c-0.28125 2.3906 1.4258 4.5586 3.8164 4.8516 2.3906 0.28906 4.5664-1.4023 4.8672-3.793l21.684-178.54 28.129-42.516 3.7617 1.5547v-0.003906c9.8164 3.9531 17.699 11.586 21.969 21.27 5.207 12.082 4.0742 25.965-3.0195 37.043l-1.457 2.2773c-0.36328 0.57031-0.58984 1.2148-0.66406 1.8867l-6.6133 61.512h0.003906c-0.16406 1.5195 0.47656 3.0156 1.6914 3.9414 31.891 24.285 49.215 63.141 45.965 103.09-0.1875 2.4062 1.6094 4.5117 4.0156 4.7031 0.11719 0.007812 0.23438 0.011719 0.35156 0.011719 2.2812-0.003907 4.1758-1.7578 4.3555-4.0273 1.1289-14.113-0.10938-28.312-3.6602-42.016l127.69 22.395 8.4727 49.965c0.34375 2.0234 2.0977 3.5039 4.1484 3.5039 1.8672-0.003906 3.5312-1.1875 4.1406-2.9531 0.24219-0.73438 0.29297-1.5195 0.15234-2.2812l-7.7773-46.637 95.785 16.801c1.1445 0.20312 2.3203-0.058594 3.2695-0.72266 0.95313-0.66797 1.5977-1.6875 1.7969-2.832l69.539-396.45c0.19922-1.1406-0.0625-2.3164-0.72656-3.2695-0.66797-0.94922-1.6836-1.5938-2.8281-1.7969l-448.15-78.605c-2.3789-0.41797-4.6484 1.1719-5.0664 3.5508l-69.539 396.45c-0.41797 2.3789 1.1719 4.6445 3.5508 5.0625l67.184 11.789 13.082 98.09h0.003906c0.28906 2.1758 2.1445 3.7969 4.3359 3.7969 1.2773 0 2.4922-0.55078 3.332-1.5117 0.83984-0.96094 1.2227-2.2383 1.0547-3.5039zm130.39-48.199-33.617-5.8984 4.0742-37.91 32.617-88.949c2.4609-6.6367 8.2773-11.457 15.254-12.645l-1.5703 95.602zm16.285-20.961-0.39062 23.746-7.1797-1.2578zm-1.9844-142.26c-3.5195 0-6.8203-1.6992-8.8711-4.5625-2.0508-2.8594-2.5977-6.5352-1.4648-9.8672l16.109-47.414c0.67187-1.9766-0.14453-4.1484-1.9492-5.1953-3.3438-1.9688-4.7812-6.0586-3.4102-9.6875 0.75-1.9609 2.25-3.5469 4.168-4.4023 1.918-0.85156 4.0977-0.90625 6.0586-0.15234 2.0625 0.79688 3.7031 2.4141 4.5234 4.4688 0.89453 2.2422 3.4375 3.3398 5.6797 2.4492 5.4766-2.1719 10.328-5.668 14.121-10.172l22.344 13.633-13.77 31.5c-1.3633 3.1211-3.6094 5.7734-6.457 7.6367-2.8477 1.8672-6.1797 2.8594-9.582 2.8633-1.7461 0-3.4844-0.26172-5.1484-0.78125l-0.042968-0.011718c-0.050782-0.015626-0.097656-0.03125-0.14844-0.046876-0.33594-0.10937-0.67188-0.22656-1-0.35937l-0.12891-0.046876c-3.7617-1.3359-7.1406-3.5703-9.8438-6.5078-1.6328-1.7812-4.3984-1.9023-6.1836-0.26953-1.7812 1.6328-1.9023 4.4023-0.26953 6.1836 3.0234 3.2852 6.6758 5.9297 10.742 7.7773l-5.1172 15.477c-1.4727 4.4688-5.6484 7.4922-10.359 7.4883zm55.242-44.332-9.0078 41.203-25.008-9.3477 2.625-7.9258c5.6758 0.79688 11.457-0.28516 16.461-3.082s8.9531-7.1562 11.246-12.41zm-42.465 62.184 41.637 2.6719 12.242 69.031c0.30078 1.6875 1.5586 3.0469 3.2188 3.4727 1.6641 0.42578 3.4219-0.15625 4.4961-1.4922l53.426-66.297 30.297 1.9453c-5.3711 4.6016-8.9688 10.938-10.172 17.91l-24.938 142.21-112.66-19.762zm85.508-18.789v-0.003906c-0.42188-2.375-2.6836-3.9648-5.0625-3.5547-3.8398 0.68359-7.7812 0.5-11.539-0.53125l-0.16406-0.042968c-0.23828-0.0625-0.47266-0.13281-0.67188-0.19141l-0.003906-0.003906c-6.9961-2.1602-10.941-9.5625-8.8281-16.574l9.8516-29.277v-0.003906c7.0312 4.4609 15.555 5.9258 23.672 4.0664-0.33203 4.2891 0.19141 8.6016 1.543 12.688 0.42969 1.3047 1.4531 2.3359 2.7578 2.7734 1.3047 0.44141 2.7422 0.23828 3.875-0.54297 2.1758-1.4805 4.9453-1.7852 7.3906-0.8125 2.4453 0.97656 4.2461 3.1016 4.8047 5.6719 0.55859 2.5742-0.19922 5.2539-2.0195 7.1562-1.8203 1.8984-4.4648 2.7695-7.0586 2.3242-0.47266-0.085937-0.9375-0.20703-1.3906-0.37109-1.2148-0.4375-2.5625-0.31641-3.6836 0.33203-1.1172 0.64844-1.8906 1.7578-2.1133 3.0312l-5.8945 33.645c-0.21484 1.2227 0.10156 2.4805 0.87109 3.4609 0.76953 0.97656 1.918 1.5781 3.1562 1.6602l9.3281 0.59766-43.52 54.004-10.18-57.453 13.523 0.86719c2.2266 0.13672 4.1992-1.4141 4.5938-3.6094l2.4609-14.016c3.5742 0.47656 7.2031 0.39844 10.754-0.22266 2.3789-0.42188 3.9688-2.6914 3.5469-5.0703zm58.262 47.547v-0.003906c1.3516-7.6992 6.707-14.09 14.047-16.77 7.3398-2.6797 15.555-1.2383 21.547 3.7773 5.9922 5.0195 8.8516 12.852 7.5 20.551l-24.945 142.21-43.094-7.5586zm40.266 74.328v-0.003906c13.715-12.754 22.863-29.66 26.039-48.117l0.42578-2.4297h0.003906c2.5234-14.371 0.48047-29.168-5.8359-42.32-6.3203-13.148-16.598-23.992-29.395-31-16.559-9.1445-27.988-25.402-30.988-44.078l-0.85156-5.4023v-0.003907c-2.0469-13.066-8.9141-24.895-19.25-33.148-10.336-8.2578-23.387-12.34-36.586-11.449l2.4023-6.25c0.60156-1.5703 0.25-3.3477-0.90234-4.5703-1.1562-1.2227-2.9102-1.6758-4.5117-1.168l-10.082 3.2109 1.1445-2.9922 0.003907 0.003906c0.58984-1.5469 0.25781-3.2969-0.85938-4.5195-1.1172-1.2227-2.832-1.707-4.4258-1.2539l-12.938 3.6758-25.977-15.789c-1.3828-0.84766-2.8398-1.5625-4.3555-2.1367-6.668-2.5625-14.082-2.3711-20.605 0.53516-6.5273 2.9023-11.629 8.2812-14.188 14.953l-12.84 33.457c-1.207 3.1055-1.4492 6.5039-0.69531 9.7539 0.75391 3.2461 2.4688 6.1914 4.9219 8.4531l-12.031 35.422-75.73-28.301c-0.25-0.09375-0.50391-0.16406-0.76562-0.21094-8.5977-1.5078-17.449 0.25781-24.809 4.957-7.3555 4.6992-12.688 11.984-14.934 20.418l-43.777 164.3-3.082-0.53906 52.914-301.64 387.83 68.031-52.914 301.64-25.855-4.5352zm-269.22-102.23-2.6758 10.035c-0.63672 2.3203-3.0234 3.6914-5.3438 3.0703-2.3242-0.61719-3.7109-2.9961-3.1094-5.3242l2.6758-10.039v0.003907c0.29688-1.1211 1.0312-2.0781 2.0352-2.6602 1.0039-0.58203 2.1953-0.73828 3.3164-0.44141 2.3359 0.625 3.7227 3.0195 3.1016 5.3555z"></path>
            </svg>
        `;
    }

    getSvgMaxHeight() {
        const windowHeight = window.innerHeight
            || document.documentElement.clientHeight
            || document.body.clientHeight;

        const imageMaxHeight = windowHeight - 290; // тут чисто хардкод под стили для css-класса mainContainer
        return imageMaxHeight;
    }

    hide() {
        this.place.innerHTML = '';
    }
}

const app = new App();
app.render();

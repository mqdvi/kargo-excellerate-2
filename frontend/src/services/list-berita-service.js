import axios from 'axios';

const base_URL = "http://localhost:8090/v1/articles";

class ArtikelService {
    getArtikel(artikel){
        return axios.get(base_URL + artikel);
    }

    createArtikel(artikel, state){
        return axios.post(base_URL + artikel, state);
    }

    getArtikelById(artikelId) {
        return axios.get(base_URL + artikelId);
    }

    updateArtikel(artikel, state) {
        return axios.put(base_URL + artikel, state);
    }

    deleteArtikel(artikelId) {
        return axios.delete(base_URL + artikelId)
    }
}

export default new ArtikelService()
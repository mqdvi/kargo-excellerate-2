import axios from 'axios';

const base_URL = "http://localhost:8090/transporters/trucks";

class TrucksService {
    getTrucks(){
        return axios.get(base_URL);
    }

    createTrucks(state){
        return axios.post(base_URL, state);
    }

    getTrucksById(trucksId) {
        return axios.get(base_URL + trucksId);
    }

    updateTrucks(state) {
        return axios.put(base_URL, state);
    }

    deleteTrucks(trucksId) {
        return axios.delete(base_URL + trucksId)
    }
}

export default new TrucksService()
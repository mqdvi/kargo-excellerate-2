import axios from 'axios';

const base_URL = "http://localhost:8090/transporters/trucks";
const truckType = "http://localhost:8090/transporters/truck-types";

class TrucksService {
    getTrucks(){
        return axios.get(base_URL);
    }
    getTrucksType(){
        return axios.get(truckType);
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
}

export default new TrucksService()
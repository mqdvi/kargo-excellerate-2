import axios from 'axios';

const base_URL = "http://localhost:8090/transporters/drivers";

class DriverService {
    getDriver(){
        return axios.get(base_URL);
    }

    createDriver(state){
        return axios.post(base_URL, state);
    }

    getDriverById(driverId) {
        return axios.get(base_URL + driverId);
    }

    updateDriver(state) {
        return axios.put(base_URL, state);
    }

    deleteDriver(driverId) {
        return axios.delete(base_URL + driverId)
    }
}

export default new DriverService()
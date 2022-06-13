import axios from 'axios'
import { backendUrl } from '../config/url'

const apiServer = axios.create({
  baseURL: backendUrl,
  withCredentials: false, // This is the default
  headers: {
    Accept: 'application/json',
    'Content-Type': 'application/json'
  }
})

export default {
  getEmployee() {
    return apiServer.get('/test')
  }
}

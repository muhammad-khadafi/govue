import axios from 'axios'
import { backendUrl } from './config/env'

export const apiServer = axios.create({
  baseURL: backendUrl,
})

export const axiosGlobalConfig = axios

export const fileUrl = backendUrl

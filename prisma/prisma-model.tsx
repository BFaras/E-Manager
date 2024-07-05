import { AnyARecord } from "dns"
import { Key } from "react"


export interface Store {
  id: string
  name: string
  userId: string
  createdAt : Date 
  updatedAt : Date
}
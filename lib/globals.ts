import { EventEmitter } from 'events'
import { Database } from 'arn/interfaces/Database'

const aerospike = require('aero-aerospike')

export const events = new EventEmitter()
export const api = require('../security/api-keys.json')
export const db: Database = aerospike.client(require('../config.json').database)
export const production = process.env.NODE_ENV === 'production'

// Yield handlers
require('./utils/yield')

// Connect to DB
db.connect().then(() => console.log('Successfully connected to database!'))
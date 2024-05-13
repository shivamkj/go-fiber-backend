// Copyright (©) 2024 - Shivam Kumar Jha - All Rights Reserved, Proprietary and confidential
// Unauthorised copying of this file, via any medium is strictly prohibited

import { generateClass, generateSection, generateUser } from './1.master_data.js'
import { generateAttendance, generateOffDay } from './2.attendance.js'

// Master Data
await generateClass()
await generateSection()
await generateUser()

// Attendance
await generateAttendance()
await generateOffDay()

process.exit(0)

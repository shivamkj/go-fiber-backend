// Copyright (©) 2024 - Shivam Kumar Jha - All Rights Reserved, Proprietary and confidential
// Unauthorised copying of this file, via any medium is strictly prohibited

import { knex, randomBool, randomDate, randomInt, dateToISOInt } from './util.js'

const start = new Date(2024, 3, 1)
const end = new Date(2024, 9, 30)

// Generate Attendance for every student
export const generateAttendance = async () => {
  const students = await knex('student').select('*')

  for (const student of students) {
    const absentCount = randomInt(0, 30)
    const attendance = []
    const dates = []

    for (let index = 0; index < absentCount; index++) {
      const isAbsent = randomBool(0.8)

      // keep dates unique to avoid voilating unique constraints
      var newDate = dateToISOInt(randomDate(start, end))
      while (dates.some((date) => date == newDate)) {
        newDate = dateToISOInt(randomDate(start, end))
      }

      // convert false to null to make easier to count aggregate
      const falseToNull = (boolean) => {
        if (boolean) return true
        else return undefined
      }
      attendance.push({
        student_id: student.id,
        date: newDate,
        section_id: student.section_id,
        is_absent: falseToNull(isAbsent),
        is_half_day: falseToNull(!isAbsent && randomBool()),
        is_late: falseToNull(!isAbsent),
      })
      dates.push(newDate)
    }

    if (attendance.length != 0) await knex('attendance').insert(attendance)
    dates.length = 0
  }
}

// Generate weekends, holidays, offdays etc.
export const generateOffDay = async () => {
  const holidays = [
    { date: new Date(2024, 0, 1), info: 'New Year Holiday' },
    { date: new Date(2024, 0, 15), info: 'Makara Sankranti' },
    { date: new Date(2024, 0, 26), info: 'Republic Day' },
    { date: new Date(2024, 1, 14), info: 'Vasant Panchami' },
    { date: new Date(2024, 2, 8), info: 'Maha Shivaratri' },
    { date: new Date(2024, 2, 25), info: 'Holi' },
    { date: new Date(2024, 2, 29), info: 'Good Friday' },
    { date: new Date(2024, 3, 9), info: 'Gudi Padwa' },
    { date: new Date(2024, 3, 10), info: 'Idul Fitr' },
    { date: new Date(2024, 3, 17), info: 'Ram Navami' },
    { date: new Date(2024, 3, 21), info: 'Mahavir Jayanti' },
    { date: new Date(2024, 4, 23), info: 'Buddha Purnima' },
    { date: new Date(2024, 7, 17), info: 'Bakrid' },
    { date: new Date(2024, 8, 7), info: 'Ganesh Chaturthi' },
    { date: new Date(2024, 9, 2), info: 'Gandhi Jayanti' },
    { date: new Date(2024, 10, 1), info: 'Diwali' },
    { date: new Date(2024, 11, 25), info: 'Christmas Day' },
    { date: new Date(2024, 11, 31), info: "New Year's Eve" },
  ]

  const sundays = getDaysBetweenDates(start, end, days.sun)
  const saturdays = getDaysBetweenDates(start, end, days.sat)

  // Offdays for all classes - Sunday & public holidays
  const offDays = []
  for (const day of holidays) {
    offDays.push({
      date: dateToISOInt(day.date),
      isholiday: true,
      description: day.info,
    })
  }
  for (const date of sundays) {
    offDays.push({ date: dateToISOInt(date), isweekend: true })
  }
  await knex('offday').insert(offDays)

  // Offdays for Saturday Off only till class 5
  const classesTill5 = await knex('class').select('*').limit(5)
  for (const cls of classesTill5) {
    const saturdayOff = []
    for (const date of saturdays) {
      saturdayOff.push({
        date: dateToISOInt(date),
        isweekend: true,
        classid: cls.id,
      })
    }
    await knex('offday').insert(saturdayOff)
  }
}

const days = { sun: 0, mon: 1, tue: 2, wed: 3, thu: 4, fri: 5, sat: 6 }

function getDaysBetweenDates(startDate, endDate, dayIndex) {
  const currentDate = new Date(startDate) // Clone to avoid modifying the original

  // Increment the current date until it's the desired day of the week
  while (currentDate.getDay() !== dayIndex) {
    currentDate.setDate(currentDate.getDate() + 1)
  }

  // Add the first date that matches the desired day of the week
  const dates = []
  if (currentDate >= startDate && currentDate <= endDate) {
    dates.push(new Date(currentDate))
  }
  // Increment the current date by one week and check if it's within the range
  while (currentDate.setDate(currentDate.getDate() + 7) <= endDate) {
    dates.push(new Date(currentDate))
  }

  return dates
}

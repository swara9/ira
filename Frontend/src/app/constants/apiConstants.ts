export const apiUrls = {
    baseUrl : 'https://iraelearning.com/api',
    updateDescription : '/instructor/course/<courseId>/description',
    createCourses : '/instructor/course',
    instrCourses : '/instructor/courses',
    uploadVideo : '/instructor/course/<courseId>/module/video',
    getAllCourses : '/courses',
    uploadQuiz : '/instructor/course/<courseId>/module/quiz',
    getCourseDetails : '/course/<courseId>',
    publishCourse : '/instructor/course/<courseId>/publish',
    studentEnroll : '/student/course/<courseId>/enroll',
    checkEnroll : '/student/course/<courseId>/enroll',
    getCourseDescription : '/instructor/course/<courseId>/description',
    getModule: '/courses/<courseId>/module/<moduleId>',
    calculateGrade: '/student/course/<courseID>/module/quiz/score'

}
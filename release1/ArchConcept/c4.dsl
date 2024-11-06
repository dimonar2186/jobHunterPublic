workspace "Name" "Description" {

    !identifiers hierarchical

    model {
        u = person "User"
        ss = softwareSystem "jobHunter" {
            //authFront = container "Authentication Web app"
            applicationsFront = container "Applications Web App"
            //pa = container "Phone Application"
            apllicationsDb = container "Applications DB" {
                tags "Database"
            }
            //usersDb = container "Users DB" {
            //    tags "Database"
            //}
            fileDb = container "File DB" {
                tags "Database"
            }
            //aut = container "Autirization service"
            fileStorage = container "File storage service" {
                file = component "Attached files" {
                    description "See classes/class_file.puml"
                }
            }

            apiGateway = container "API gateway" "nginx"

            //integration = container "Ext systems integration Service"
            applicationsBack = container "JobHunter Back Service" {
                vacancy = component "Vacancy" {
                    description "See classes/class_vacancy.puml"
                }

                //vacancyResponse = component "Response to the vacancy"
                jobSearchingProcess = component "jobSearchingProcess" {
                    description "See classes/class_jobSearchingProcess.puml"
                }

                employer = component "Employer" {
                    description "See classes/class_employer.puml"
                }

                applicationStage = component "Stage of the application" {
                    description "See classes/class_applicationStage.puml"
                }

                HRManager = component "Human resources managers" {
                    description "See classes/class_HRManager.puml"
                }

                HRcontact = component "HR's contacts" {
                    description "See classes/class_HRContact.puml"
                }

                //notes = component "Plain text to remember important things"
                offer = component "Offer details" {
                    description "See classes/class_offer.puml"
                }

                salaryRange = component "Salary range" {
                    description "See classes/class_salatyRange.puml"
                }

                vacancyHRManager = component "Vacancy to HRManagers connection" {
                    description "See classes/class_vacancyHRManager.puml"
                }
            }
        }
        //ess = softwareSystem "HeadHunter" {
        //   tags "existingSoftwareSystem"
        //    url "http://hh.ru"
        //}
        //google = softwareSystem "Google SSO" {
        //    tags "existingSoftwareSystem"
        //}

        u -> ss.applicationsFront "Uses"
        //u -> ss.authFront "Authenticates"

        //ss.authFront -> ss.aut "Authorize users"
        //ss.pa -> ss.aut "Authorize users"


        ss.applicationsFront -> ss.apiGateway "Manage job searching process" 
        ss.apiGateway -> ss.applicationsBack "Create, delete, update of vacancies, jobSearchingProcesses etc."
        ss.apiGateway -> ss.applicationsBack.jobSearchingProcess "Starts and finishes job searching process"
        ss.apiGateway -> ss.applicationsBack.vacancy "Create, read, delete and update"
        //ss.applicationsFront -> ss.applicationsBack.vacancyResponse "Create, read and update"
        ss.apiGateway -> ss.applicationsBack.applicationStage "Add and remove stages"
        ss.apiGateway -> ss.applicationsBack.employer "Choose employer"
        ss.apiGateway -> ss.applicationsBack.HRManager "Add and update HR"
        //ss.applicationsFront -> ss.applicationsBack.notes  "Add and update notes"
        //ss.applicationsBack -> ss.aut "Check permissions"
        //ss.applicationsBack.vacancy -> ss.applicationsBack.vacancyResponse "Continues with"
        ss.applicationsBack.HRManager -> ss.applicationsBack.vacancy "Relates to"
        ss.applicationsBack.applicationStage -> ss.applicationsBack.vacancy "Part of"
        ss.applicationsBack.employer -> ss.applicationsBack.vacancy "Associated with"
        ss.applicationsBack.vacancy -> ss.applicationsBack.jobSearchingProcess "Starts"
        //ss.applicationsBack.notes -> ss.applicationsBack.applicationStage "Relates to"
        //ss.applicationsBack.notes -> ss.applicationsBack.hr "From"
        ss.applicationsBack.HRManager -> ss.applicationsBack.HRcontact "Reachable via"
        ss.applicationsBack.vacancy -> ss.applicationsBack.offer "Results with"

        ss.applicationsBack -> ss.apllicationsDb "Reads from and writes to"
        //ss.aut -> ss.usersDb "Reads from and writes to"
        
        ss.apiGateway -> ss.fileStorage "Upload, delete, download files"
        ss.apiGateway -> ss.fileStorage.file "Upload, remove, download"
        //ss.applicationsFront -> ss.fileStorage.vacancyDescription "Upload, remove, download vacancy description"
        //ss.fileStorage -> ss.aut "Check permissions"
        //ss.fileStorage.cv -> ss.applicationsBack.vacancyResponse "Describes response"
        //ss.fileStorage.vacancyDescription -> ss.applicationsBack.vacancyResponse "Describes vacancy"
        
        //ss.aut -> google "Authorize users" 
        //ss.pa -> ss.wa "Sends to and gets from"
        //ss.applicationsFront -> ss.integration "Get vacancies"
        //ss.integration -> ess "Gets vacancies"
        //ss.integration -> ss.aut "Check permissions"
        //ss.integration -> ss.applicationsBack "Uses"
        ss.fileStorage -> ss.fileDb "Reads and writes to"
        ss.fileStorage.file -> ss.fileDb "Reads and writes to"

        
    }

    views {
        systemContext ss "Diagram1" {
            include *
            autolayout lr
        }

        container ss "jobHunterContainerDiagram" {
            include *
            autolayout lr
        }
        

        component ss.applicationsBack "jobHunter_applicationsBack_ComponentDiagram" {
            include *
            autolayout lr
        }

        component ss.fileStorage "jobHunter_fileStorage_ComponentDiagram" {
            include *
            autolayout lr
        }

        styles {
            element "Element" {
                color #ffffff
            }
            element "Person" {
                background #757475
                shape person
            }
            element "Software System" {
                background #8723d9
            }
            element "Container" {
                background #9a28f8
            }
            element "Database" {
                shape cylinder
            }
            element "existingSoftwareSystem" {
                background #48414E
            }
        }
    }

    configuration {
        scope softwaresystem
    }

}
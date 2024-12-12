*THIS IS A BUSINESS IDEA AND FOR THAT REASON, THE COMPLEMENTORY PROJECT (E-COMMERCE) DOESN'T HAVE THE COMPLETED VERSION ON GITHUB AND .ENV FILE IS HIDDEN. 
I HAVE ONLY PUT E-MANAGER IN DISPLAY TO SHOWCASE MY CODING SKILLS.STILL, I WILL PROVIDE SOME PICTURES OF OF HOW E-COMMERCE LOOKS LIKE" 

Project Overview
This project allows users to create multiple stores, add products, and manage store-related operations. Each store can have multiple products, but before adding a product, a billboard, category, and size must be created. Once a product is added, it becomes available for purchase in the store, and the dashboard is updated to provide analytics for the user, including total revenue and sales metrics.

Authentication
The project uses Clerk for authentication. Upon running the application, users are redirected to the sign-in/sign-up page depending on whether they have an existing account.

Why Clerk?
Clerk is a good choice if the primary requirement is authentication only.
Firebase is better suited for projects that require a broader set of features beyond authentication, such as real-time databases, cloud functions, and storage.
Backend Architecture
The backend follows a three-layered architecture to keep the code modular and maintainable. It uses the repository pattern to perform CRUD operations for entities like stores, products, billboards, categories, and sizes.

Repository Pattern: This pattern is used to handle database interactions, keeping the code clean and organized. However, the decision to avoid using an ORM (Object-Relational Mapping) was made to gain a deeper understanding of raw SQL queries and database management. While this approach has been educational, it led to repetitive code, which could have been avoided with an ORM.

Takeaway: Using an ORM like Prisma or GORM might have been more efficient, as it simplifies database interactions and reduces the risk of SQL injection. Nonetheless, working directly with SQL has improved my understanding of query construction and security practices.

JWT Authentication
I am using JWT tokens for secure authentication between the frontend (Clerk) and the backend. JWT tokens are issued by Clerk and passed to the backend for verification before any user-related action is performed. This process has been valuable in deepening my understanding of JWT-based authentication workflows.

Key Learnings
Clerk vs. Firebase: Firebase is better for projects that require more than just authentication, whereas Clerk is more lightweight and focused on simplifying authentication alone.
Avoiding ORM: While beneficial for learning raw SQL, avoiding an ORM resulted in more repetitive code and potential security risks like SQL injection.
JWT Implementation: Integrating JWT authentication with Clerk improved security and broadened my knowledge of token-based authentication.

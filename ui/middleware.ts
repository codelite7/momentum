// with cookie has to be last
// export default stackMiddlewares([withLogging, withAuth]);
import { authkitMiddleware } from "@workos-inc/authkit-nextjs";

export default authkitMiddleware();

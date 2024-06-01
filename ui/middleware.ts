import { withLogging } from "@/middlewares/withLogging";
import { stackMiddlewares } from "@/middlewares/stackMiddlewares";
import { withAuth } from "@/middlewares/withAuth";

// with cookie has to be last
export default stackMiddlewares([withLogging, withAuth]);

import { parse } from 'url';
import micri, {
  IncomingMessage,
  ServerResponse,
  Router,
  send
} from 'micri';
import hello from './hello';
import csp from './csp-report'

const { router, on, otherwise } = Router;
const PORT = process.env.PORT || 3000

const parsePath = (req: IncomingMessage): string => parse(req.url || '/').path || '/';
console.log("Starting at http://localhost:" + PORT)
micri(router(
  on.get((req: IncomingMessage) => parsePath(req) === '/', hello),
  on.get((req: IncomingMessage) => parsePath(req) === '/_/csp', csp),
  otherwise((_req: IncomingMessage, res: ServerResponse) => send(res, 400, 'Method Not Accepted'))))
  .listen(PORT);
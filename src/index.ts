import { parse } from 'url';
import micri, {
  IncomingMessage,
  ServerResponse,
  Router,
  send
} from 'micri';
import hello from './hello';
import csp from './csp'

const { router, on, otherwise } = Router;

const parsePath = (req: IncomingMessage): string => parse(req.url || '/').path || '/';

micri(router(
  on.get((req: IncomingMessage) => parsePath(req) === '/', hello),
  on.get((req: IncomingMessage) => parsePath(req) === '/csp', csp),
  otherwise((_req: IncomingMessage, res: ServerResponse) => send(res, 400, 'Method Not Accepted'))))
  .listen(3000);
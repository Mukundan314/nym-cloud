/* tslint:disable */
/* eslint-disable */
/**
*/
export function set_panic_hook(): void;
/**
*/
export class NymClient {
  free(): void;
/**
* @param {string} validator_server
*/
  constructor(validator_server: string);
/**
* @param {Function} on_message
*/
  set_on_message(on_message: Function): void;
/**
* @param {Function} on_connect
*/
  set_on_gateway_connect(on_connect: Function): void;
/**
* @returns {string}
*/
  self_address(): string;
/**
* @returns {Promise<NymClient>}
*/
  initial_setup(): Promise<NymClient>;
/**
* @param {string} message
* @param {string} recipient
* @returns {Promise<NymClient>}
*/
  send_message(message: string, recipient: string): Promise<NymClient>;
/**
* @returns {Promise<NymClient>}
*/
  get_and_update_topology(): Promise<NymClient>;
}

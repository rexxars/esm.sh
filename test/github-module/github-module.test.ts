import { assertEquals } from "https://deno.land/std@0.180.0/testing/asserts.ts";

import sfMeta from "http://localhost:8080/gh/superfluid-finance/metadata";

Deno.test("github module", async () => {
  const network = sfMeta.getNetworkByName("eth-goerli");
  assertEquals(network?.name, "eth-goerli");
});

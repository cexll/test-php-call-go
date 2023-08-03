<?php

namespace App\Controller;

use Mix\Vega\Context;
use Swlib\SaberGM;

class Hello
{
    private \Mix\Grpc\Client $cli;
    public function __construct()
    {
        $this->connGrpc();
    }

    private function connGrpc()
    {
//        $this->cli->close();
        $this->cli = new \Mix\Grpc\Client('127.0.0.1', 8081, false, 6000);
    }

    /**
     * @param Context $ctx
     */
    public function http(Context $ctx)
    {
        try {
            $resp = SaberGM::get('http://localhost:8082/v1/demo/hello?id=123');

            $body = $resp->getBody()->getContents();

            $ctx->JSON(200, \json_decode($body));
        } catch (\Throwable $e) {
            $ctx->string(500, '');
        }
    }

    public function grpc(Context $ctx)
    {
        try {
            $hello = new \Demo\Api\V1\HelloClient($this->cli);
            $req = new \Demo\Api\V1\HelloReq();
            $req->setId("123123123123");
            $gCtx = new \Mix\Grpc\Context();
            $res = $hello->sayHello($gCtx, $req);
            $ctx->JSON(200, [
                'id' => $res->getId(),
                'message' => $res->getMessage(),
            ]);
        } catch (\Throwable $e) {
//            var_dump((string)$e);
            $this->connGrpc();
        }
    }
}

<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/merchant-basic/invoice.proto

namespace GPBMetadata\Proto\MerchantBasic;

class Invoice
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        $pool->internalAddGeneratedFile(
            '
�
"proto/merchant-basic/invoice.proto
CreateInvoiceRequest
action (	
invoice_data (	"B
CreateInvoiceResponse

error_code (

        , true);

        static::$is_initialized = true;
    }
}

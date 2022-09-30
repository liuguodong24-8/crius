<?php

declare(strict_types=1);

namespace Omy\MicroServers\Client;

use Omy\Crius\BasicServer\AllAreasResponse;
use Omy\Crius\BasicServer\AreaRequest;
use Omy\Crius\BasicServer\AreasResponse;
use Omy\Crius\BasicServer\EmptyRequest;
use Omy\Crius\BasicServer\GetAllAreasByLevelRequest;
use Omy\Crius\BasicServer\GetAllAreasByLevelResponse;
use Omy\Crius\BasicServer\GetSubAreasRequest;
use Omy\Crius\BasicServer\GetSubAreasResponse;
use Omy\Crius\BasicServer\PhoneAreaResponse;
use Omy\Crius\BasicServer\ProvinceAreasResponse;
use Omy\Crius\BasicServer\Response;
use Omy\Crius\BasicServer\ShowAreaRequest;
use Omy\Crius\BasicServer\ShowAreaResponse;
use Omy\Crius\BasicServer\UpdateAreaRequest;

class BasicClient extends GrpcClient
{
    private const BASE_PATH = '/basic.BasicService/';

    public const GET_ALL_AREAS_BY_LEVEL = 'GetAllAreasByLevel';

    public const GET_AREAS = 'GetAreas';

    public const GET_AREAS_BY_PROVINCE = 'GetAreasByProvince';

    public const GET_PHONE_AREAS = 'GetPhoneAreas';

    public const GET_ALL_AREAS = 'GetAllAreas';

    public const SHOW_AREA = 'ShowArea';

    public const UPDATE_AREA = 'UpdateArea';

    public const GET_SUB_AREA = 'GetSubAreas';


    protected array $actionMap = [
        self::GET_ALL_AREAS_BY_LEVEL => [
            GetAllAreasByLevelRequest::class,
            GetAllAreasByLevelResponse::class,
        ],
        self::GET_AREAS => [
            AreaRequest::class,
            AreasResponse::class
        ],
        self::GET_AREAS_BY_PROVINCE => [
            AreaRequest::class,
            ProvinceAreasResponse::class
        ],
        self::GET_PHONE_AREAS => [
            EmptyRequest::class,
            PhoneAreaResponse::class
        ],
        self::GET_ALL_AREAS => [
            EmptyRequest::class,
            AllAreasResponse::class
        ],
        self::SHOW_AREA => [
            ShowAreaRequest::class,
            ShowAreaResponse::class
        ],
        self::UPDATE_AREA => [
            UpdateAreaRequest::class,
            Response::class
        ],
        self::GET_SUB_AREA => [
            GetSubAreasRequest::class,
            GetSubAreasResponse::class
        ]
    ];

    protected function getBasePath()
    {
        return self::BASE_PATH;
    }
}

package handler

// func getPayloadData(ctx *fiber.Ctx, dto struct{}) (interface{}, error, builder.Response) {
// 	lg := config.AppConfigutarion.GetLogger()
// 	lgFields := []zap.Field{zap.String("Method", "getPayloadData")}
// 	lg.Info(fmt.Sprintf(utils.REGISTER_HANDLER, utils.STARTED))

// 	err := ctx.BodyParser(&dto)
// 	if err != nil {
// 		lgFields = append(lgFields, zap.Any(utils.ERROR, err))
// 		lg.Error(utils.BODY_PARSER_ERROR, lgFields...)
// 		responseBuilder := builder.Response{
// 			Ctx: ctx,
// 			ErrorRes: builder.ErrorResponse{
// 				Message: utils.BadRequest,
// 				Code:    utils.StatusCode[utils.BadRequest],
// 				Error:   fmt.Sprintf("%v", err),
// 			},
// 			Status: utils.StatusCode[utils.BadRequest],
// 		}
// 		lg.Info(fmt.Sprintf(utils.REGISTER_HANDLER, utils.END))
// 		return nil, err, responseBuilder
// 	}
// 	lg.Info(fmt.Sprintf(utils.REGISTER_HANDLER, utils.END_WITH_ERROR))
// 	return dto, err, builder.Response{}
// }

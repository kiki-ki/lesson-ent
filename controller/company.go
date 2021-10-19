package controller

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/kiki-ki/lesson-ent/controller/request"
	"github.com/kiki-ki/lesson-ent/database"
	"github.com/kiki-ki/lesson-ent/ent/user"
	"github.com/kiki-ki/lesson-ent/util"
)

func NewCompanyController(dbc *database.EntClient) CompanyController {
	return &companyController{
		dbc: dbc,
		ctx: context.Background(),
	}
}

type CompanyController interface {
	Show(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
	IndexUsers(http.ResponseWriter, *http.Request)
	CreateWithUser(http.ResponseWriter, *http.Request)
}

type companyController struct {
	dbc *database.EntClient
	ctx context.Context
}

func (c *companyController) Show(w http.ResponseWriter, r *http.Request) {
	cId, err := strconv.Atoi(chi.URLParam(r, "companyId"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}
	company, err := c.dbc.Company.Get(c.ctx, cId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, company)
}

func (c *companyController) Update(w http.ResponseWriter, r *http.Request) {
	cId, err := strconv.Atoi(chi.URLParam(r, "companyId"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}
	company, err := c.dbc.Company.Get(c.ctx, cId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, err.Error())
		return
	}
	var req request.CompanyUpdateReq
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, err.Error())
		return
	}
	company, err = company.Update().SetName(req.Name).Save(c.ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, nil)
		return
	}
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, company)
}

func (c *companyController) Delete(w http.ResponseWriter, r *http.Request) {
	cId, err := strconv.Atoi(chi.URLParam(r, "companyId"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}
	company, err := c.dbc.Company.Get(c.ctx, cId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		render.JSON(w, r, err.Error())
		return
	}
	err = c.dbc.Company.DeleteOne(company).Exec(c.ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, fmt.Sprintf("id=%d is deleted", cId))
}

func (c *companyController) IndexUsers(w http.ResponseWriter, r *http.Request) {
	cId, err := strconv.Atoi(chi.URLParam(r, "companyId"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}
	company, err := c.dbc.Company.Get(c.ctx, cId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		render.JSON(w, r, err.Error())
		return
	}
	users, err := company.QueryUsers().All(c.ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, users)
}

func (c *companyController) CreateWithUser(w http.ResponseWriter, r *http.Request) {
	var req request.CompanyCreateWithUserReq
	err := render.DecodeJSON(r.Body, &req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, err.Error())
		return
	}
	tx, err := c.dbc.Tx(c.ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, err.Error())
		return
	}
	company, err := tx.Company.
		Create().
		SetName(req.CompanyName).
		Save(c.ctx)
	if err != nil {
		err = util.Rollback(tx, err)
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, err.Error())
		return
	}
	user, err := tx.User.Create().
		SetCompany(company).
		SetName(req.UserName).
		SetEmail(req.UserEmail).
		SetRole(user.RoleAdmin).
		SetComment(req.UserComment).
		Save(c.ctx)
	if err != nil {
		err = util.Rollback(tx, err)
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, err.Error())
		return
	}
	if err := tx.Commit(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, map[string]interface{}{
		"company": company,
		"user":    user,
	})
}

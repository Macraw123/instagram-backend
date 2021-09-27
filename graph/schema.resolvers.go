package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"Go/graph/generated"
	"Go/graph/model"
	"context"
	"errors"
	"fmt"
	"log"
	"net/smtp"
	"strconv"
)

func (r *mutationResolver) AddUser(ctx context.Context, userEmail string, name string, username string, password string) (*model.User, error) {
	myBio := new(string)
	*myBio = "Hello Love!"

	myV := new(bool)
	*myV = false

	myG := new(bool)
	*myG = false

	myT := new(string)
	*myT = "piu12" + username

	newUser := &model.User{
		Name:     name,
		Username: username,
		Email:    userEmail,
		Password: password,
		Google:   myG,
		Verified: myV,
		Bio:      myBio,
		Token:    myT,
	}

	nameChecker := new(model.User)
	errName := r.DB.Model(nameChecker).Where("\"username\"=?", username).Select()

	emailChecker := new(model.User)
	errEmail := r.DB.Model(emailChecker).Where("\"email\"=?", userEmail).Select()

	if errEmail == nil {
		return nil, errors.New("Email is already exist!")
	} else if userEmail == "" {
		return nil, errors.New("Email cannot be empty!")
	} else if name == "" {
		return nil, errors.New("Full name cannot be empty!")
	} else if errName == nil {
		return nil, errors.New("Username is already exist!")
	} else if username == "" {
		return nil, errors.New("Username cannot be empty!")
	} else if password == "" {
		return nil, errors.New("Password cannot be empty!")
	}
	_, err := r.DB.Model(newUser).Column("name", "username", "email", "password",  "google", "verified",  "bio", "token").Insert()
	if err != nil {
		return nil, err
	}
	giveEmail("Please click this link: http://localhost:3000/verified?email="+userEmail, userEmail)

	return newUser, nil
}

func (r *mutationResolver) ChangePassword(ctx context.Context, userEmail string, password string) (*model.User, error) {
	if password == "" {
		return nil, errors.New("Password cannot be empty!")
	} else if userEmail == "" {
		return nil, errors.New("Email cannot be empty!")
	}

	user := new(model.User)
	_, err := r.DB.Model(user).Set("password=?", password).Where("email=?", userEmail).Update()

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *mutationResolver) VerifiedStatus(ctx context.Context, userEmail string) (*model.User, error) {
	user := new(model.User)
	_, err := r.DB.Model(user).Set("verified=?", true).Where("email=?", userEmail).Update()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *mutationResolver) AddSearchHistory(ctx context.Context, email string, word string) (*model.User, error) {
	history := model.Searchhistory{
		Word:  word,
		Email: email,
	}
	_, err := r.DB.Model(&history).Column("word", "email").Insert()

	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *mutationResolver) ViewStory(ctx context.Context, storyid int, email string) (*model.Story, error) {
	view := model.Viewedstory{
		Storyid:   storyid,
		Viewed:    "yes",
		Useremail: email,
	}

	_, err := r.DB.Model(&view).Column("storyid", "viewed", "useremail").Insert()

	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *mutationResolver) AddCommentToPost(ctx context.Context, postID int, userEmail string, comment string) (*model.User, error) {
	comments := new(model.Comment)
	users := new(model.User)

	erruser := r.DB.Model(&users).Column("user.*").Where("email=?", userEmail).First()

	if erruser != nil {
		return nil, erruser
	}
	comments.Comment = comment
	comments.Userid, _ = strconv.Atoi(users.ID)

	comments.Postid = postID

	_, err := r.DB.Model(&comments).Column("comment", "userid", "postid").Insert()
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (r *mutationResolver) LikePost(ctx context.Context, postID int, userEmail string) (*model.User, error) {
	likes := model.Like{
		Postid:    postID,
		Useremail: userEmail,
	}

	_, err := r.DB.Model(&likes).Column("postid", "useremail").Insert()
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *mutationResolver) UnlikePost(ctx context.Context, postID int, userEmail string) (*model.User, error) {
	likes := new(model.Like)

	_, err := r.DB.Model(&likes).Column("like.*").Where("postid=?", postID).Where("useremail=?", userEmail).Delete()
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *mutationResolver) SavePost(ctx context.Context, postID int, email string) (*model.User, error) {
	save := model.Save{
		Postid:    postID,
		Useremail: email,
	}
	_, err := r.DB.Model(&save).Column("postid", "useremail").Insert()
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *mutationResolver) UnsavePost(ctx context.Context, postID int, email string) (*model.User, error) {
	saves := new(model.User)

	_, err := r.DB.Model(&saves).Column("save.*").Where("postid=?", postID).Where("useremail=?", email).Delete()
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *mutationResolver) AddReply(ctx context.Context, commentID int, userEmail string, comment string) (*model.User, error) {
	reply := model.Reply{
		Commentid: commentID,
		Useremail: userEmail,
		Comment:   comment,
	}

	_, err := r.DB.Model(&reply).Column("commentid", "useremail", "comment").Insert()
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *mutationResolver) DeletePostByID(ctx context.Context, postID int) (*model.Post, error) {
	post := new(model.Post)
	_, err := r.DB.Model(&post).Where("id=? ", postID).Delete()
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (r *mutationResolver) UpdatePostCaption(ctx context.Context, postID int, caption string) (*model.Post, error) {
	post := new(model.Post)
	_, err := r.DB.Model(&post).Set("caption=?", caption).Where("id=?", postID).Update()

	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *mutationResolver) DeleteCommentByCommentID(ctx context.Context, commentID int) (*model.Comment, error) {
	comment := new(model.Comment)
	_, err := r.DB.Model(&comment).Where("id=? ", commentID).Delete()
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *mutationResolver) DeleteReplyByReplyID(ctx context.Context, replyID int) (*model.Reply, error) {
	reply := new(model.Reply)
	_, err := r.DB.Model(&reply).Where("id=?", replyID).Delete()
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *mutationResolver) SendMessage(ctx context.Context, toemail string, fromemail string, created string, message string, image *string, video *string) (*model.Message, error) {
	messages := model.Message{
		Toemail:   toemail,
		Fromemail: fromemail,
		Created:   created,
		Message:   message,
		Image:     image,
		Video:     video,
	}
	_, err := r.DB.Model(&messages).Column("toemail", "fromemail", "created", "message", "image", "video").Insert()
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *mutationResolver) UnsendMessage(ctx context.Context, id int) (*model.Message, error) {
	messsage := new(model.Message)
	_, err := r.DB.Model(&messsage).Where("id=?", id).Delete()

	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *mutationResolver) CreateChatRoom(ctx context.Context, owneremail string, recipientemail string) (*model.Chatroom, error) {
	room := model.Chatroom{
		Owneremail:     owneremail,
		Recipientemail: recipientemail,
	}

	_, err := r.DB.Model(&room).Column("owneremail", "recipientemail").Insert()

	if err != nil {
		return nil, err
	}
	return &room, nil
}

func (r *mutationResolver) DeleteChatRoom(ctx context.Context, roomID int) (*model.Chatroom, error) {
	room := new(model.Chatroom)

	_, err := r.DB.Model(&room).Where("id=?", roomID).Delete()
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *mutationResolver) DeleteStoryByID(ctx context.Context, storyID int) (*model.Story, error) {
	story := new(model.Story)

	_, err := r.DB.Model(&story).Where("id=?", storyID).Delete()
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *mutationResolver) LikeCommentByCommentID(ctx context.Context, commentid int, email string) (*model.Commentlike, error) {
	commentLike := model.Commentlike{
		Commentid: commentid,
		Useremail: email,
	}
	_, err := r.DB.Model(&commentLike).Column("commentid", "useremail").Insert()

	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *mutationResolver) UnlikeCommentByCommentID(ctx context.Context, commentid int, email string) (*model.Commentlike, error) {
	commentLike := new(model.Like)

	_, err := r.DB.Model(&commentLike).Where("commentid=? ", commentid).Where("useremail=?", email).Delete()

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (r *mutationResolver) CreatePost(ctx context.Context, caption string, image string, video string, userid int, created string) (*model.Post, error) {
	post := model.Post{
		Caption: caption,
		Userid:  userid,
		Created: created,
		Video:   &video,
	}

	r.DB.Model(&post).Column("caption", "userid", "created", "video").Insert()

	r.DB.Model(&post).Column("post.*").Last()

	t, _ := strconv.Atoi(post.ID)
	imageModel := model.Image{
		Postid:    t,
		Imagelink: image,
	}

	r.DB.Model(&imageModel).Column("postid", "imagelink").Insert()

	return &post, nil
}

func (r *queryResolver) Login(ctx context.Context, userEmail string, password string) (*model.User, error) {
	user := new(model.User)
	err := r.DB.Model(user).Where(" \"email\"=? AND \"password\"=?", userEmail, password).Select()

	if err != nil {
		return nil, errors.New("Invalid email or password!")
	} else if *user.Verified == false {
		return nil, errors.New("Email not verified!")
	}
	return user, nil
}

func (r *queryResolver) SendEmail(ctx context.Context, userEmail string, typeArg string) (*model.User, error) {
	if typeArg == "forgot_password" {
		giveEmail("Please click this link: http://localhost:3000/verified_password?email="+userEmail, userEmail)
	} else if typeArg == "verify" {
		giveEmail("Please click this link: http://localhost:3000/verified?email="+userEmail, userEmail)
	}
	user := new(model.User)
	return user, nil
}

func (r *queryResolver) ForgotPassword(ctx context.Context, userEmail string) (*model.User, error) {
	emailChecker := new(model.User)
	errEmail := r.DB.Model(emailChecker).Where("email=?", userEmail).Select()

	if userEmail == "" {
		return nil, errors.New("Email cannot be empty!")
	} else if errEmail != nil {
		return nil, errors.New("Email doesn't exist")
	}
	giveEmail("Please click this link: http://localhost:3000/verified_password?email="+userEmail, userEmail)
	return emailChecker, nil
}

func (r *queryResolver) SearchUsers(ctx context.Context, username string) ([]*model.User, error) {
	var user []*model.User
	err := r.DB.Model(&user).Column("user.*").Where("username LIKE ?", username+"%").Select()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *queryResolver) Comments(ctx context.Context, postID int) ([]*model.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Stories(ctx context.Context, userID int) ([]*model.Story, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Post(ctx context.Context, postID int) (*model.Post, error) {
	var post []*model.Post
	var pictures []*model.Image
	user := new(model.User)
	var comments []*model.Comment
	var users []*model.User

	errCheckComment := r.DB.Model(&comments).Column("comment.*").Select()
	if errCheckComment != nil {
		return nil, errors.New("Invalid Selection Comments")
	}
	errCheckUsers := r.DB.Model(&users).Column("user.*").Select()
	if errCheckUsers != nil {
		return nil, errors.New("Invalid Selection Users")
	}

	for _, c := range comments {
		for _, u := range users {
			if strconv.Itoa(c.Userid) == u.ID {
				c.User = append(c.User, u)
			}
		}
	}

	errPost := r.DB.Model(&post).Column("post.*").Where("id=?", postID).First()

	if errPost != nil {
		return nil, errPost
	}

	r.DB.Model(&user).Column("user.*").Where("id=?", post[0].Userid).First()

	err := r.DB.Model(&pictures).Column("image.*").Select()

	if err != nil {
		return nil, errors.New("Invalid Selection of Image")
	}

	for _, i := range pictures {
		if strconv.Itoa(i.Postid) == post[0].ID {
			post[0].Image = append(post[0].Image, i)
		}
	}

	for _, c := range comments {
		if strconv.Itoa(c.Postid) == post[0].ID {
			post[0].Comment = append(post[0].Comment, c)
		}
	}

	post[0].User = user
	return post[0], nil
}

func (r *queryResolver) GetFollowersByID(ctx context.Context, userID int) ([]*model.User, error) {
	var followers []*model.Follower
	user := new(model.User)
	var users []*model.User
	var myfollowers []*model.User

	errID := r.DB.Model(&user).Column("user.*").Where("id=?", userID).First()
	if errID != nil {
		return nil, errors.New("Unknown ID")
	}

	errFollowerId := r.DB.Model(&followers).Column("follower.*").Where("userid = ?", user.ID).Select()
	if errFollowerId != nil {
		return nil, errors.New("Unknown Follower ID")
	}

	errSelection := r.DB.Model(&users).Column("user.*").Select()
	if errSelection != nil {
		return nil, errors.New("Invalid Selection of Columns")
	}

	for _, f := range followers {
		for _, u := range users {
			if strconv.Itoa(f.Followerid) == u.ID {
				myfollowers = append(myfollowers, u)
			}
		}
	}

	return myfollowers, nil
}

func (r *queryResolver) GetFollowingsByID(ctx context.Context, userID int) ([]*model.User, error) {
	var followers []*model.Follower
	var user model.User
	var users []*model.User
	var myfollowings []*model.User

	r.DB.Model(&user).Column("user.*").Where("id=?", userID).First()

	r.DB.Model(&followers).Column("follower.*").Where("followerid = ?", user.ID).Select()

	r.DB.Model(&users).Column("user.*").Select()

	for _, f := range followers {
		for _, u := range users {
			if strconv.Itoa(f.Userid) == u.ID {
				myfollowings = append(myfollowings, u)
			}
		}
	}
	return myfollowings, nil
}

func (r *queryResolver) GetPostByID(ctx context.Context, userID int) ([]*model.Post, error) {
	var posts []*model.Post
	var pictures []*model.Image
	var following []*model.User
	var result []*model.Post
	var comments []*model.Comment
	var users []*model.User
	var as []model.User

	r.DB.Model(&comments).Select()
	r.DB.Model(&as).Select()

	fmt.Println(r.DB.Model(&as))

	for _, c := range comments {
		for _, u := range users {
			if strconv.Itoa(c.Userid) == u.ID {
				c.User = append(c.User, u)
			}
		}
	}
	following, _ = r.GetFollowingsByID(ctx, userID)
	r.DB.Model(&posts).Column("post.*").Select()

	err := r.DB.Model(&pictures).Column("image.*").Select()

	if err != nil {
		return nil, nil
	}

	for _, p := range posts {
		for _, i := range pictures {
			if strconv.Itoa(i.Postid) == p.ID {
				p.Image = append(p.Image, i)
			}
		}
		for _, c := range comments {
			if strconv.Itoa(c.Postid) == p.ID {
				p.Comment = append(p.Comment, c)
			}
		}
		var pid, _ = strconv.Atoi(p.ID)

		var postLikes []*model.Like

		r.DB.Model(&postLikes).Column("like.*").Where("postid=?", pid).Select()

		var likeCount = len(postLikes)
		p.Like = &likeCount
	}

	for _, f := range following {
		for _, p := range posts {
			if strconv.Itoa(p.Userid) == f.ID {
				p.User = f
				result = append(result, p)
			}
		}
	}
	return result, nil
}

func (r *queryResolver) GetPictureByPostID(ctx context.Context, postID int) ([]*model.Image, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetCommentByPostID(ctx context.Context, postID *int) ([]*model.Comment, error) {
	var comments []*model.Comment
	var users []*model.User
	r.DB.Model(&comments).Column("comment.*").Where("postid=?", postID).Select()
	r.DB.Model(&users).Column("user.*").Select()

	for _, c := range comments {
		for _, u := range users {
			if strconv.Itoa(c.Userid) == u.ID {
				c.User = append(c.User, u)
			}
		}
	}

	return comments, nil
}

func (r *queryResolver) GetStoryByUserID(ctx context.Context, userid int) ([]*model.Story, error) {
	var stories []*model.Story
	var user []*model.User

	err := r.DB.Model(&user).
		Column("user.*").Where("id=?", userid).
		Select()

	r.DB.Model(&stories).Column("story.*").Where("userid=?", userid).Select()

	for _, sr := range stories {
		sr.User = user[0]
	}

	if err != nil {
		return nil, err
	}

	return stories, nil
}

func (r *queryResolver) GetMainStoryByID(ctx context.Context, userid int) ([]*model.Story, error) {
	var following []*model.User
	var result []*model.Story
	var story []*model.Story

	following, _ = r.GetFollowingsByID(ctx, userid)

	r.DB.Model(&story).Column("story.*").Select()

	for _, fw := range following {
		for _, sr := range story {
			if strconv.Itoa(sr.Userid) == fw.ID {
				sr.User = fw
				result = append(result, sr)
			}
		}
	}

	return result, nil
}

func (r *queryResolver) GetMainStoryUserByID(ctx context.Context, userid int) ([]*model.User, error) {
	var following []*model.User
	var result []*model.User
	var story []*model.Story

	following, _ = r.GetFollowingsByID(ctx, userid)

	r.DB.Model(&story).Column("story.userid").Distinct().Select()

	for _, fw := range following {
		for _, sr := range story {
			if strconv.Itoa(sr.Userid) == fw.ID {
				result = append(result, fw)
			}
		}
	}

	return result, nil
}

func (r *queryResolver) GetMyOwnStory(ctx context.Context, email string) ([]*model.Story, error) {
	var mystory []*model.Story
	var user model.User

	r.DB.Model(&user).Column("user.*").Where("email=?", email).First()

	r.DB.Model(&mystory).Column("story.*").Where("userid=?", user.ID).Select()

	return mystory, nil
}

func (r *queryResolver) GetCommentLike(ctx context.Context, commentid int) ([]*model.Commentlike, error) {
	var commentlike []*model.Commentlike

	r.DB.Model(&commentlike).Where("commentid=?", commentid).Select()

	return commentlike, nil
}

func (r *queryResolver) GetSavedPost(ctx context.Context, email string) ([]*model.Post, error) {
	var saves []*model.Save
	var posts []*model.Post
	var res []*model.Post
	var images []*model.Image

	err := r.DB.Model(&saves).Column("save.*").Where("useremail =?", email).Select()

	r.DB.Model(&posts).Column("post.*").Select()
	r.DB.Model(&images).Column("image.*").Select()

	if err != nil {
		return nil, err
	}

	for _, s := range saves {
		for _, p := range posts {
			if strconv.Itoa(s.Postid) == p.ID {
				res = append(res, p)
			}
		}
	}

	for _, p := range posts {
		for _, i := range images {
			if p.ID == strconv.Itoa(i.Postid) {
				p.Image = append(p.Image, i)
			}
		}
	}

	return res, nil
}

func (r *queryResolver) GetRepliesByCommentID(ctx context.Context, commentid int) ([]*model.Reply, error) {
	var replies []*model.Reply
	var users []*model.User

	err := r.DB.Model(&replies).Column("reply.*").Where("commentid=?", commentid).Select()

	if err != nil {
		return nil, err
	}

	r.DB.Model(&users).Column("user.*").Select()

	for _, r := range replies {
		for _, u := range users {
			if r.Useremail == u.Email {
				r.User = u
			}
		}
	}

	return replies, nil
}

func (r *queryResolver) GetMessageByToEmail(ctx context.Context, toEmail string, fromEmail string) ([]*model.Message, error) {
	var messages []*model.Message
	r.DB.Model(&messages).Column("message.*").Where("toemail=?", toEmail).Where("fromemail=?", fromEmail).WhereOr("toemail=?", fromEmail).Where("fromemail=?", toEmail).Order("created ASC").Select()

	return messages, nil
}

func (r *queryResolver) GetChatHistory(ctx context.Context, myEmail string) ([]*model.User, error) {
	var users []*model.User
	var messages1 []*model.Message
	var messages2 []*model.Message
	var results []*model.User

	r.DB.Model(&messages1).Column("message.*").Where("toemail=?", myEmail).Distinct().Select()
	r.DB.Model(&messages2).Column("message.*").Where("fromemail=?", myEmail).Distinct().Select()

	r.DB.Model(&users).Column("user.*").Select()

	r.DB.Model(&results).Column("user.*").Where("email=?", "anis@ymail.com").First()

	for _, m := range messages1 {
		for _, u := range users {
			for _, r := range results {
				if u.Email == m.Fromemail && r.Email != u.Email {
					results = append(results, u)
					break
				}

			}
		}
	}

	for _, m := range messages2 {
		for _, u := range users {
			for _, r := range results {
				if u.Email == m.Toemail && r.Email != u.Email {
					results = append(results, u)
					break

				}
			}
		}
	}

	return results, nil
}

func (r *queryResolver) GetMyChats(ctx context.Context, myEmail string) ([]*model.Chatroom, error) {
	var cr []*model.Chatroom
	var users []*model.User

	err := r.DB.Model(&cr).Column("chatroom.*").Where("owneremail=?", myEmail).Select()
	r.DB.Model(&users).Column("user.*").Select()

	for _, c := range cr {
		for _, u := range users {
			if u.Email == c.Recipientemail {
				c.User = u
			}
		}
	}

	if err != nil {
		return nil, err
	}

	return cr, nil
}

func (r *queryResolver) GetSearchHistory(ctx context.Context, email string) ([]*model.Searchhistory, error) {
	var hist []*model.Searchhistory

	err := r.DB.Model(&hist).Column("searchhistory.*").Where("email=?", email).Select()

	if err != nil {
		return nil, err
	}
	return hist, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func giveEmail(message string, email string) {
	from := "jsentosa743@gmail.com"
	fromEl := fmt.Sprintf("From: <%s>\n", "jsentosa743@gmail.com")
	password := "jokosentosa123"
	to := []string{email}
	toEl := fmt.Sprintf("To: <%s>\n", email)
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	subject := "Subject: This is Rub Bish\n\n"
	body := message + "\n"

	msg := fromEl + toEl + subject + body
	messages := []byte(msg)

	// Create authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Send actual message
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, messages)
	if err != nil {
		log.Fatal(err)
	}
}
